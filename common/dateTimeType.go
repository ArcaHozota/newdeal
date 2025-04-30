package common

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// DateTime は「YYYY-MM-DD HH:MM:SS.ffffff」固定の日時型
type DateTime struct{ time.Time }

// MarshalJSON: 値 → `"2006-01-02 15:04:05.123456"`
func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte(`""`), nil // 空文字列で NULL 表現
	}
	str := d.Time.Format(DateTimeLayout)
	return []byte(`"` + str + `"`), nil
}

// UnmarshalJSON: `"1994-12-03 10:11:12.000000"` → 値
func (d *DateTime) UnmarshalJSON(b []byte) error {
	trim := bytes.Trim(b, `"`)
	if len(trim) == 0 { // 空はゼロ値
		*d = DateTime{}
		return nil
	}
	t, err := time.ParseInLocation(DateTimeLayout, string(trim), time.Local) // ← UTC にしたい場合は time.UTC
	if err != nil {
		return fmt.Errorf("DateTime: %w", err)
	}
	d.Time = t
	return nil
}

// Value: driver.Valuer (INSERT/UPDATE)
func (d DateTime) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil // NULL
	}
	return d.Time.Format(DateTimeLayout), nil
}

// Scan: sql.Scanner (SELECT)
func (d *DateTime) Scan(src any) error {
	switch v := src.(type) {
	case time.Time:
		d.Time = v
	case []byte:
		t, err := time.ParseInLocation(DateTimeLayout, string(v), time.Local)
		if err != nil {
			return err
		}
		d.Time = t
	case string:
		t, err := time.ParseInLocation(DateTimeLayout, v, time.Local)
		if err != nil {
			return err
		}
		d.Time = t
	case nil:
		*d = DateTime{}
	default:
		return fmt.Errorf("DateTime: unsupported Scan type %T", v)
	}
	return nil
}
