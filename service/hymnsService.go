package service

import (
	"newdeal/common"
	"newdeal/models"
	"newdeal/pojos"
	"newdeal/repository"

	"github.com/samber/lo"
)

func CountHymnsByKeyword(keyword string) (uint32, error) {
	kennsu, err := repository.CountHymnsByKeyword(keyword)
	return kennsu, err
}

func GetHymnsByKeyword(keyword string, pageNum int) ([]pojos.HymnDTO, error) {
	hymns, err := repository.PaginationHymns(keyword, pageNum, int(common.DefaultPageSize))
	return lo.Map(hymns, func(hm models.Hymn, _ int) pojos.HymnDTO {
		return pojos.HymnDTO{
			ID:          hm.Id,
			NameJP:      hm.NameJp,
			NameKR:      hm.NameKr,
			Serif:       hm.Serif,
			Link:        hm.Link,
			Score:       nil,
			Biko:        common.EmptyString,
			UpdatedUser: hm.UpdatedUser,
			UpdatedTime: hm.UpdatedTime,
			LineNumber:  0,
		}
	}), err
}
