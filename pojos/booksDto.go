package pojos

// BookDTO は Java の BookDto に相当する Go 構造体。
type BookDTO struct {
	ID     string `json:"id"`     // ID
	Name   string `json:"name"`   // 名称
	NameJP string `json:"nameJp"` // 日本語名称
}

// ChapterDTO は Java の ChapterDto に相当する Go 構造体。
type ChapterDTO struct {
	ID     string `json:"id"`     // ID
	Name   string `json:"name"`   // 名称
	NameJP string `json:"nameJp"` // 日本語名称
	BookID string `json:"bookId"` // 書別ID
}

// PhraseDTO は Java の PhraseDto に相当する Go 構造体。
type PhraseDTO struct {
	ID        string `json:"id"`        // ID
	Name      string `json:"name"`      // 名称
	TextEN    string `json:"textEn"`    // 内容
	TextJP    string `json:"textJp"`    // 日本語内容
	ChapterID string `json:"chapterId"` // 書別ID
}
