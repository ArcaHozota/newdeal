package pojos

// ChapterDTO は Java の ChapterDto に相当する Go 構造体。
type ChapterDTO struct {
	ID     string `json:"id"`     // ID
	Name   string `json:"name"`   // 名称
	NameJP string `json:"nameJp"` // 日本語名称
	BookID string `json:"bookId"` // 書別ID
}
