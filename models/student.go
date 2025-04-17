package models

import "time"

type Student struct {
	Id           int64     `gorm:"primaryKey"`
	LoginAccount string    `gorm:"uniqueIndex;not null"`
	Password     string    `gorm:"not null"`
	Username     string    `gorm:"not null"`
	DateOfBirth  string    `gorm:"not null"`
	Email        string    `gorm:"uniqueIndex"`
	UpdatedTime  time.Time `gorm:"not null"`
	VisibleFlg   bool      `gorm:"not null"`
}

// func (Student) TableName() string {
// 	return "students"
// }
