package service

import (
	"context"
	"fmt"
	"newdeal/common"
	"newdeal/ent"
	"newdeal/ent/book"
	"newdeal/ent/chapter"
	"newdeal/ent/phrase"
	"newdeal/pojos"
	"strconv"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/samber/lo"
)

func GetBooks() ([]pojos.BookDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	books, err := EntCore.Book.Query().
		Order(book.ByID(sql.OrderAsc())).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return map2BookDTOs(books), nil
}

func GetChaptersByBookId(bookId int16) ([]pojos.ChapterDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	chapters, err := EntCore.Chapter.Query().
		Where(
			chapter.BookIDEQ(bookId),
		).
		Order(chapter.ByID(sql.OrderAsc())).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return map2ChapterDTOs(chapters), nil
}

func PhraseInfoStorage(phraseDto pojos.PhraseDTO) (string, error) {
	id, err := strconv.Atoi(phraseDto.ID)
	if err != nil {
		return common.EmptyString, err
	}
	chapterId, err := strconv.Atoi(phraseDto.ChapterID)
	if err != nil {
		return common.EmptyString, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	queriedCh, err := EntCore.Chapter.Query().Where(
		chapter.IDEQ(int32(chapterId)),
	).Only(ctx)
	if err != nil {
		return common.EmptyString, err
	}
	changeLineMark := false
	textEn := phraseDto.TextEN
	if strings.HasSuffix(textEn, "#") {
		changeLineMark = true
		textEn = strings.ReplaceAll(textEn, "#", common.EmptyString)
	}
	phraseId := int64(chapterId*100 + id)
	count, _ := EntCore.Phrase.Query().
		Where(
			phrase.IDEQ(phraseId),
		).
		Count(ctx)
	if count != 0 {
		err = EntCore.Phrase.UpdateOneID(phraseId).
			SetName(fmt.Sprintf("%s:%d", queriedCh.Name, id)).
			SetTextEn(textEn).
			SetTextJp(phraseDto.TextJP).
			SetChapterID(queriedCh.ID).
			SetChangeLine(changeLineMark).
			Exec(ctx)
		if err != nil {
			return common.EmptyString, err
		}
		return common.UpsertedMsg, nil
	}
	err = EntCore.Phrase.Create().
		SetID(phraseId).
		SetName(fmt.Sprintf("%s:%d", queriedCh.Name, id)).
		SetTextEn(textEn).
		SetTextJp(phraseDto.TextJP).
		SetChapterID(queriedCh.ID).
		SetChangeLine(changeLineMark).
		Exec(ctx)
	if err != nil {
		return common.EmptyString, err
	}
	return common.UpsertedMsg, nil
}

// BookDTOへマップする
func map2BookDTOs(books []*ent.Book) []pojos.BookDTO {
	return lo.Map(books, func(bk *ent.Book, _ int) pojos.BookDTO {
		return pojos.BookDTO{
			ID:     strconv.Itoa(int(bk.ID)),
			Name:   bk.Name,
			NameJP: bk.NameJp,
		}
	})
}

// ChapterDTOへマップする
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
