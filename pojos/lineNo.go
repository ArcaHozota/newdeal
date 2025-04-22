package pojos

import (
	"encoding/json"
	"fmt"
)

// LineNumber は行番号を識別する列挙型。
type LineNumber int

const (
	// 順序を Java と同じ 2,1,3,5 に合わせたいなら個別に指定
	BURGUNDY LineNumber = 2
	CADMIUM  LineNumber = 1
	NAPLES   LineNumber = 3
	SNOWY    LineNumber = 5
)

var lineNumberToName = map[LineNumber]string{
	BURGUNDY: "BURGUNDY",
	CADMIUM:  "CADMIUM",
	NAPLES:   "NAPLES",
	SNOWY:    "SNOWY",
}

var nameToLineNumber = map[string]LineNumber{
	"BURGUNDY": BURGUNDY,
	"CADMIUM":  CADMIUM,
	"NAPLES":   NAPLES,
	"SNOWY":    SNOWY,
}

// Stringer インタフェースを実装して fmt.Printf などで名前を表示
func (ln LineNumber) String() string {
	if s, ok := lineNumberToName[ln]; ok {
		return s
	}
	return fmt.Sprintf("LineNumber(%d)", ln)
}

// JSON で "BURGUNDY" のように出力したい場合
func (ln LineNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(ln.String())
}

func (ln *LineNumber) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if v, ok := nameToLineNumber[s]; ok {
		*ln = v
		return nil
	}
	return fmt.Errorf("unknown LineNumber %q", s)
}

// Java の getLineNo() 相当：整数値を返すヘルパー
func (ln LineNumber) Int() int {
	return int(ln)
}
