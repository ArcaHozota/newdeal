package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "host=52.199.163.241 user=postgres password=postgres dbname=nasb1995 port=5432 sslmode=disable TimeZone=Asia/Seoul"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
}
