package service

import (
	"context"
	"log"
	"newdeal/ent"

	_ "github.com/lib/pq"
)

var EntCore *ent.Client
var err error

func init() {
	// DSNを定義する
	dsn := "postgres://postgres:postgres@52.199.163.241:5432/nasb1995?sslmode=disable"
	// entを配置する
	EntCore, err = ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer EntCore.Close()
	// Run the auto migration tool.
	if err := EntCore.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
