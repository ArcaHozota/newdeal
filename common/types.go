package common

import (
	"strings"
	"time"
)

// time.Time をラップ
type Date struct{ time.Time }

// UnmarshalJSON: "1994-12-03" → Date
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// MarshalJSON: Date → "1994-12-03"
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format("2006-01-02") + `"`), nil
}

func NewDate(t time.Time) Date {
	// たとえば必ずローカルタイムゾーン 0:00 に揃える
	y, m, d := t.Date()
	return Date{Time: time.Date(y, m, d, 0, 0, 0, 0, nil)}
}
