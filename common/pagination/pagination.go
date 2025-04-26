package pagination

import (
	"errors"
	"fmt"
)

// Pagination 与 Java 版字段一一对应。
// T 表示记录元素类型；可以是你的实体模型，也可以是 DTO。
type Pagination[T any] struct {
	Records       []T   `json:"records"`       // 当前页数据
	PageNum       int   `json:"pageNum"`       // 当前页（从 1 开始）
	PageSize      int   `json:"pageSize"`      // 每页大小
	TotalPages    int   `json:"totalPages"`    // 总页数
	TotalRecords  int   `json:"totalRecords"`  // 总记录数
	HasPrevPage   bool  `json:"hasPrevPage"`   // 是否有上一页
	HasNextPage   bool  `json:"hasNextPage"`   // 是否有下一页
	PrevPage      int   `json:"prevPage"`      // 上一页 (0 表示不存在)
	NextPage      int   `json:"nextPage"`      // 下一页 (0 表示不存在)
	NavigatePages int   `json:"navigatePages"` // 最大导航页数
	NaviFirstPage int   `json:"naviFirstPage"` // 导航条第一个页码
	NaviLastPage  int   `json:"naviLastPage"`  // 导航条最后一个页码
	NavigateNums  []int `json:"navigateNos"`   // 实际导航页码切片
}

// NewPagination 创建分页对象；pageNum 从 1 开始。
func NewPagination[T any](
	records []T,
	totalRecords int,
	pageNum int,
	pageSize int,
	navigatePages int,
) (*Pagination[T], error) {

	if pageSize <= 0 || pageNum <= 0 {
		return nil, errors.New("pageSize 和 pageNum 必须 > 0")
	}
	if records == nil {
		return nil, errors.New("records 不能为 nil")
	}

	p := &Pagination[T]{
		Records:       records,
		PageNum:       pageNum,
		PageSize:      pageSize,
		TotalRecords:  totalRecords,
		NavigatePages: navigatePages,
	}

	// 总页数
	p.TotalPages = (totalRecords + pageSize - 1) / pageSize

	// 边界检查，防止越界页码
	if p.PageNum > p.TotalPages {
		p.PageNum = p.TotalPages
	}
	if p.TotalPages == 0 { // 没有记录
		p.PageNum = 1
	}

	// 上下页
	p.HasPrevPage = p.PageNum > 1
	p.HasNextPage = p.PageNum < p.TotalPages
	if p.HasPrevPage {
		p.PrevPage = p.PageNum - 1
	}
	if p.HasNextPage {
		p.NextPage = p.PageNum + 1
	}

	// 计算导航条页码
	p.calcNavigate()

	return p, nil
}

// HasContent 判断当前页是否有数据
func (p *Pagination[T]) HasContent() bool {
	return len(p.Records) > 0
}

// fmt.Stringer
func (p *Pagination[T]) String() string {
	return fmt.Sprintf(
		"Pagination{page=%d/%d, size=%d, total=%d}",
		p.PageNum, p.TotalPages, p.PageSize, p.TotalRecords,
	)
}

// ------- 私有辅助 --------

func (p *Pagination[T]) calcNavigate() {
	if p.TotalPages == 0 {
		p.NavigateNums = []int{}
		return
	}

	// 导航条不够长时直接全量输出
	if p.TotalPages <= p.NavigatePages {
		p.NavigateNums = make([]int, p.TotalPages)
		for i := 0; i < p.TotalPages; i++ {
			p.NavigateNums[i] = i + 1
		}
		p.NaviFirstPage = 1
		p.NaviLastPage = p.TotalPages
		return
	}

	half := p.NavigatePages / 2
	start := p.PageNum - half
	end := p.PageNum + half

	if start < 1 {
		start = 1
		end = p.NavigatePages
	} else if end > p.TotalPages {
		end = p.TotalPages
		start = p.TotalPages - p.NavigatePages + 1
	}

	p.NavigateNums = make([]int, 0, p.NavigatePages)
	for i := start; i <= end; i++ {
		p.NavigateNums = append(p.NavigateNums, i)
	}
	p.NaviFirstPage = start
	p.NaviLastPage = end
}

// ------------------- GORM 辅助 Scope -------------------

// PaginateScope 用于 GORM 链式调用：db.Scopes(PaginateScope(page, size))
// func PaginateScope(page, size int) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if size <= 0 {
// 			return db
// 		}
// 		offset := (page - 1) * size
// 		return db.Offset(offset).Limit(size)
// 	}
// }
