package models

import "time"

type HymnsWork struct {
	WorkId      int64 `gorm:"primaryKey"`
	Score       []byte
	NameJpRa    string    `gorm:"column:name_jp_rational"`
	UpdatedTime time.Time `gorm:"not null"`
	Biko        string
}

func (HymnsWork) TableName() string {
	return "hymns_work"
}
