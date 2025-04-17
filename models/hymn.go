package models

import "time"

type Hymn struct {
	Id          int64     `gorm:"primaryKey"`
	NameJp      string    `gorm:"uniqueIndex;not null"`
	NameKr      string    `gorm:"uniqueIndex;not null"`
	Link        string    `gorm:"uniqueIndex"`
	UpdatedTime time.Time `gorm:"not null"`
	UpdatedUser int64     `gorm:"not null"`
	Serif       string    `gorm:"not null"`
	VisibleFlg  bool      `gorm:"not null"`
}

// func (Hymn) TableName() string {
// 	return "hymns"
// }
