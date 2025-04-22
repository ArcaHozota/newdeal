package service

import (
	"context"
	"log"
	"math/rand"
	"newdeal/common"
	"newdeal/common/tools"
	"newdeal/ent"
	"newdeal/ent/hymn"
	"newdeal/ent/hymnswork"
	"newdeal/pojos"
	"strconv"
	"strings"
	"time"

	"slices"

	"github.com/samber/lo"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func CountHymnsAll() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return EntCore.Hymn.Query().
		Where(
			hymn.VisibleFlg(true),
		).Count(ctx)
}

func GetHymnById(id int64) (*ent.Hymn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return EntCore.Hymn.Query().Where(
		hymn.VisibleFlg(true),
		hymn.ID(id),
	).Only(ctx)
}

func CountHymnsByKeyword(keyword string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return EntCore.Hymn.Query().
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
	return map2DTOs(hymns, pojos.LineNumber(5)), err
}

func GetHymnsRandomFive(keyword string) ([]pojos.HymnDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for _, strange := range common.StrangeArray {
		if strings.Contains(strings.ToLower(keyword), strange) || len(keyword) >= 100 {
			hymns, err := EntCore.Hymn.Query().
				Where(hymn.VisibleFlg(true)).
				Order(hymn.ByID()).
				Limit(int(common.DefaultPageSize)).All(ctx)
			log.Printf("怪しいキーワード: %v\n", keyword)
			return map2DTOs(hymns, pojos.LineNumber(5)), err
		}
	}
	if keyword == common.EmptyString {
		hymns, err := EntCore.Hymn.Query().
			Where(hymn.VisibleFlg(true)).All(ctx)
		hymnDtos := map2DTOs(hymns, pojos.LineNumber(5))
		return randomFiveLoop2(hymnDtos), err
	}
	hymns, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpEQ(keyword),
				hymn.NameKrEQ(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains("["+keyword+"]"),
				),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}
	withName := map2DTOs(hymns, pojos.LineNumber(1))
	hymnDtos := slices.Clone(withName)
	withNameIds := lo.Map(withName, func(hm pojos.HymnDTO, _ int) int64 {
		parseInt, err := strconv.ParseInt(hm.ID, 10, 64)
		if err != nil {
			return 0
		}
		return parseInt
	})
	hymns2, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}
	hymns2 = lo.Filter(hymns2, func(a *ent.Hymn, _ int) bool {
		return !lo.Contains(withNameIds, a.ID)
	})
	withNameLike := map2DTOs(hymns2, pojos.LineNumber(2))
	hymnDtos = append(hymnDtos, withNameLike...)
	if len(hymnDtos) >= int(common.DefaultPageSize) {
		return randomFiveLoop2(hymnDtos), err
	}
	withNameLikeIds := lo.Map(withNameLike, func(hm pojos.HymnDTO, _ int) int64 {
		parseInt, err := strconv.ParseInt(hm.ID, 10, 64)
		if err != nil {
			return 0
		}
		return parseInt
	})
	keyword = tools.GetDetailKeyword(keyword)
	hymns3, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true),
			hymn.Or(
				hymn.NameJpContains(keyword),
				hymn.NameKrContains(keyword),
				hymn.SerifContains(keyword),
				hymn.HasToWorkWith(
					hymnswork.NameJpRationalContains(keyword),
				),
			),
		).All(ctx)
	if err != nil {
		return nil, err
	}
	hymns3 = lo.Filter(hymns3, func(a *ent.Hymn, _ int) bool {
		return !lo.Contains(withNameIds, a.ID) && !lo.Contains(withNameLikeIds, a.ID)
	})
	withSerifLike := map2DTOs(hymns3, pojos.LineNumber(3))
	hymnDtos = append(hymnDtos, withSerifLike...)
	if len(hymnDtos) >= int(common.DefaultPageSize) {
		return randomFiveLoop2(hymnDtos), err
	}
	withSerifLikeIds := lo.Map(withSerifLike, func(hm pojos.HymnDTO, _ int) int64 {
		parseInt, err := strconv.ParseInt(hm.ID, 10, 64)
		if err != nil {
			return 0
		}
		return parseInt
	})
	hymns4, err := EntCore.Hymn.Query().
		Where(hymn.VisibleFlg(true)).All(ctx)
	if err != nil {
		return nil, err
	}
	hymns4 = lo.Filter(hymns4, func(a *ent.Hymn, _ int) bool {
		return !lo.Contains(withNameIds, a.ID) && !lo.Contains(withNameLikeIds, a.ID) && !lo.Contains(withSerifLikeIds, a.ID)
	})
	filteredRecords := map2DTOs(hymns4, pojos.LineNumber(5))
	hymnDtos = randomFiveLoop(hymnDtos, filteredRecords)
	return hymnDtos, err
}

func randomFiveLoop(hymnsRecords, totalRecords []pojos.HymnDTO) []pojos.HymnDTO {
	idSet := make(map[string]struct{})
	for _, h := range hymnsRecords {
		idSet[h.ID] = struct{}{}
	}
	var filteredRecords []pojos.HymnDTO
	for _, item := range totalRecords {
		if _, exists := idSet[item.ID]; !exists {
			filteredRecords = append(filteredRecords, item)
		}
	}
	concernList1 := slices.Clone(hymnsRecords) // コピー
	if len(hymnsRecords) < int(common.DefaultPageSize) {
		sagaku := int(common.DefaultPageSize) - len(hymnsRecords)
		for range sagaku {
			index := random.Intn(len(filteredRecords))
			concernList1 = append(concernList1, filteredRecords[index])
		}
	}
	concernList2 := distinctHymnDtos(concernList1)
	if len(concernList2) == int(common.DefaultPageSize) {
		return concernList2
	}
	return randomFiveLoop(concernList2, filteredRecords)
}

func randomFiveLoop2(hymnsRecords []pojos.HymnDTO) []pojos.HymnDTO {
	var concernList1 []pojos.HymnDTO
	for range int(common.DefaultPageSize) {
		index := random.Intn(len(hymnsRecords))
		concernList1 = append(concernList1, hymnsRecords[index])
	}
	concernList2 := distinctHymnDtos(concernList1)
	if len(concernList2) == int(common.DefaultPageSize) {
		return concernList2
	}
	return randomFiveLoop(concernList2, hymnsRecords)
}

func distinctHymnDtos(input []pojos.HymnDTO) []pojos.HymnDTO {
	seen := make(map[string]struct{})
	var result []pojos.HymnDTO
	for _, h := range input {
		if _, exists := seen[h.ID]; !exists {
			seen[h.ID] = struct{}{}
			result = append(result, h)
		}
	}
	return result
}

func map2DTOs(hymns []*ent.Hymn, lineNo pojos.LineNumber) []pojos.HymnDTO {
	return lo.Map(hymns, func(hm *ent.Hymn, _ int) pojos.HymnDTO {
		return pojos.HymnDTO{
			ID:          strconv.FormatInt(hm.ID, 10),
			NameJP:      hm.NameJp,
			NameKR:      hm.NameKr,
			Serif:       hm.Serif,
			Link:        hm.Link,
			Score:       nil,
			Biko:        common.EmptyString,
			UpdatedUser: strconv.FormatInt(hm.UpdatedUser, 10),
			UpdatedTime: hm.UpdatedTime,
			LineNumber:  lineNo,
		}
	})
}
