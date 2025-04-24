package pojos

// PhraseDTO は Java の PhraseDto に相当する Go 構造体。
type PhraseDTO struct {
	ID        string `json:"id"`        // ID
	Name      string `json:"name"`      // 名称
	TextEN    string `json:"textEn"`    // 内容
	TextJP    string `json:"textJp"`    // 日本語内容
	ChapterID string `json:"chapterId"` // 書別ID
}
