package service

import (
	"context"
	"fmt"
	"newdeal/common"
	"newdeal/ent"
	"newdeal/ent/chapter"
	"newdeal/ent/phrase"
	"newdeal/pojos"
	"strconv"
	"strings"
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
	if strings.HasSuffix(phraseDto.TextEN, "#") {
		changeLineMark = true
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
			SetTextEn(phraseDto.TextEN).
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
		SetTextEn(phraseDto.TextEN).
		SetTextJp(phraseDto.TextJP).
		SetChapterID(queriedCh.ID).
		SetChangeLine(changeLineMark).
		Exec(ctx)
	if err != nil {
		return common.EmptyString, err
	}
	return common.UpsertedMsg, nil
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
