package service

import (
	"context"
	"newdeal/ent"
	"newdeal/ent/chapter"
	"newdeal/pojos"
	"strconv"
	"time"

	"github.com/samber/lo"
)

func GetChaptersByBookId(bookId int16) ([]pojos.ChapterDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	chapters, err := EntCore.Chapter.Query().Where(
		chapter.BookIDEQ(bookId),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	return map2ChapterDTOs(chapters), nil
}

// DTOへマップする
func map2ChapterDTOs(chapters []*ent.Chapter) []pojos.ChapterDTO {
	return lo.Map(chapters, func(ch *ent.Chapter, _ int) pojos.ChapterDTO {
		return pojos.ChapterDTO{
			ID:     strconv.Itoa(int(ch.ID)),
			Name:   ch.Name,
			NameJP: ch.NameJp,
			BookID: strconv.Itoa(int(ch.BookID)),
		}
	})
}
