package common

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// Date time.Time をラップ
type Date struct{ time.Time }

// MarshalJSON : Date → "1994-12-03"
func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte(`""`), nil
	}
	str := d.Time.In(jst).Format(DateLayout) // ★ 画面は JST
	return []byte(`"` + str + `"`), nil
}

// Scan --- sql.Scanner (DB から読み込む) ---
func (d *Date) Scan(src interface{}) error {
	switch v := src.(type) {
	case time.Time: // DATE 型を driver が time.Time で返す RDB もある
		d.Time = v
		return nil
	case []byte:
		t, err := time.Parse(DateLayout, string(v))
		if err != nil {
			return err
		}
		d.Time = t
		return nil
	case string:
		t, err := time.Parse(DateLayout, v)
		if err != nil {
			return err
		}
		d.Time = t
		return nil
	case nil:
		*d = Date{}
		return nil
	}
	return fmt.Errorf("dto.Date: unsupported Scan, storing %T", src)
}

// UnmarshalJSON : "1994-12-03" → Date
func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(bytes.Trim(b, `"`))
	if s == "" {
		*d = Date{}
		return nil
	}
	t, err := time.ParseInLocation(DateLayout, s, jst) // JST として解釈
	if err != nil {
		return fmt.Errorf("DateTime: %w", err)
	}
	d.Time = t.UTC() // ★ 内部は UTC で保持
	return nil
}

// Value --- driver.Valuer (DB へ書き込む) ---
func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil // NULL 保存を許可
	}
	return d.Time.Format(DateLayout), nil // DATE 列には文字列でも OK
}

// Date へ追加
func (d Date) String() string {
	if d.Time.IsZero() {
		return EmptyString
	}
	return d.Time.In(jst).Format(DateLayout) // ← JST で成形
}
