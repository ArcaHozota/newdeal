package common

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

var jst = func() *time.Location { // アプリ起動時に一度だけロード
	loc, _ := time.LoadLocation("Asia/Tokyo")
	return loc
}()

// DateTime は「YYYY-MM-DD HH:MM:SS.ffffff」固定の JST 文字列を
// 内部では必ず UTC で保持するためのラッパー型
type DateTime struct{ time.Time }

///////////////////////////////////////////////////////////////////////////////
// JSON

// MarshalJSON: UTC → JST 文字列
func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte(`""`), nil
	}
	str := d.Time.In(jst).Format(DateTimeLayout) // ★ 画面は JST
	return []byte(`"` + str + `"`), nil
}

// UnmarshalJSON: JST 文字列 → UTC
func (d *DateTime) UnmarshalJSON(b []byte) error {
	s := string(bytes.Trim(b, `"`))
	if s == "" {
		*d = DateTime{}
		return nil
	}
	t, err := time.ParseInLocation(DateTimeLayout, s, jst) // JST として解釈
	if err != nil {
		return fmt.Errorf("DateTime: %w", err)
	}
	d.Time = t.UTC() // ★ 内部は UTC で保持
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// DB (timestamp[tz] 列を想定)

func (d DateTime) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil
	}
	// driver は time.Time を渡せば自動で UTC かオフセット付きで書き込む
	return d.Time.UTC(), nil // ★ 保存は UTC
}

func (d *DateTime) Scan(src any) error {
	switch v := src.(type) {
	case time.Time: // driver が既に UTC ロケーションを持たせてくれる
		d.Time = v.UTC()
	case string:
		t, err := time.Parse(time.RFC3339Nano, v) // timestamptz → RFC形式文字列になる例も
		if err != nil {
			return err
		}
		d.Time = t.UTC()
	case []byte:
		t, err := time.Parse(time.RFC3339Nano, string(v))
		if err != nil {
			return err
		}
		d.Time = t.UTC()
	case nil:
		*d = DateTime{}
	default:
		return fmt.Errorf("DateTime: unsupported Scan type %T", src)
	}
	return nil
}
