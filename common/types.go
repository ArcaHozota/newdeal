package common

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// Date time.Time をラップ
type Date struct{ time.Time }

// UnmarshalJSON : "1994-12-03" → Date
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalJSON : Date → "1994-12-03"
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(DateLayout) + `"`), nil
}

// --- sql.Scanner (DB から読み込む) ---
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

// --- driver.Valuer (DB へ書き込む) ---
func (d Date) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return nil, nil // NULL 保存を許可
	}
	return d.Time.Format(DateLayout), nil // DATE 列には文字列でも OK
}
