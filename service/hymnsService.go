package service

import (
	"context"
	"log"
	"math/rand"
	"newdeal/common"
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
	}
	if keyword == common.EmptyString {
		hymns, err := EntCore.Hymn.Query().
			Where(hymn.VisibleFlg(true)).All(ctx)
		hymnDtos := lo.Map(hymns, func(hm *ent.Hymn, _ int) pojos.HymnDTO {
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
		})
		return randomFiveLoop2(hymnDtos), err
	}
	return nil, err
}

func randomFiveLoop(hymnsRecords, totalRecords []pojos.HymnDTO) []pojos.HymnDTO {
	idSet := make(map[string]struct{})
	for _, h := range hymnsRecords {
		idSet[strconv.FormatInt(h.ID, 10)] = struct{}{}
	}
	var filteredRecords []pojos.HymnDTO
	for _, item := range totalRecords {
		if _, exists := idSet[strconv.FormatInt(item.ID, 10)]; !exists {
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
		if _, exists := seen[strconv.FormatInt(h.ID, 10)]; !exists {
			seen[strconv.FormatInt(h.ID, 10)] = struct{}{}
			result = append(result, h)
		}
	}
	return result
}
