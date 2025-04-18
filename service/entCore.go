package service

import (
	"context"
	"log"
	"newdeal/ent"

	_ "github.com/lib/pq"
)

var EntClient *ent.Client
var err error

func init() {
	// 1. DSN 推荐用 URI 格式
	dsn := "postgres://postgres:postgres@52.199.163.241:5432/nasb1995?sslmode=disable"
	// entを配置する
	EntClient, err = ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer EntClient.Close()
	// Run the auto migration tool.
	if err := EntClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
