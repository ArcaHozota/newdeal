package service

import (
	"newdeal/common"
	"newdeal/models"
	"newdeal/pojos"
	"newdeal/repository"
	"time"

	"github.com/samber/lo"
)

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
			UpdatedUser: 0,
			UpdatedTime: time.Date(1900, 1, 1, 12, 0, 0, 0, time.Local),
			LineNumber:  0,
		}
	}), err
}
