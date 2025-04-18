package service

import (
	"context"
	"newdeal/common"
	"newdeal/ent"
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/pojos"
	"time"

	"github.com/samber/lo"
)

func CountHymnsAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	kennsu, err := EntCore.Hymn.Query().
		Where(
			hymn.VisibleFlg(true),
		).Count(ctx)
	return kennsu, err
}

func CountHymnsByKeyword(keyword string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	kennsu, err := EntCore.Hymn.Query().
		Where(
			hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).Count(ctx)
	return kennsu, err
}

func GetHymnsByKeyword(keyword string, pageNum int) ([]pojos.HymnDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	offset := (pageNum - 1) * int(common.DefaultPageSize)
	hymns, err := EntCore.Hymn.Query().
		Where(
			hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).
		Order(hymn.ByID()).
		Limit(int(common.DefaultPageSize)).
		Offset(offset).All(ctx)
	return lo.Map(hymns, func(hm *ent.Hymn, _ int) pojos.HymnDTO {
		return pojos.HymnDTO{
			ID:          hm.ID,
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
