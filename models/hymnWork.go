package models

import "time"

type HymnWork struct {
	WorkId      int64 `gorm:"primaryKey"`
	Score       []byte
	NameJpRa    string    `gorm:"column:name_jp_rational"`
	UpdatedTime time.Time `gorm:"not null"`
	Biko        string
}

func (HymnWork) TableName() string {
	return "hymns_work"
}
