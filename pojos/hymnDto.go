package pojos

import (
	"newdeal/common"
)

// HymnDTO は Java の HymnDto に相当する Go 構造体。
type HymnDTO struct {
	ID          string          `json:"id"`          // ID
	NameJP      string          `json:"nameJp"`      // 日本語名称
	NameKR      string          `json:"nameKr"`      // 韓国語名称
	Serif       string          `json:"serif"`       // セリフ
	Link        string          `json:"link"`        // ビデオリンク
	Score       []byte          `json:"score"`       // 楽譜 (バイナリ)
	Biko        string          `json:"biko"`        // 備考
	UpdatedUser string          `json:"updatedUser"` // 更新者
	UpdatedTime common.DateTime `json:"updatedTime"` // 更新時間 (ISO 形式など推奨)
	LineNumber  LineNumber      `json:"lineNumber"`  // LINENUMBER
}
