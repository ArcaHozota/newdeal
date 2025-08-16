package service

import (
	"context"
	"log"
	"newdeal/ent"

	_ "github.com/lib/pq"
)

var EntClient *ent.Client
var err error

func InitEntClient() {
	// DSNを定義する
	dsn := "postgres://postgres:postgres@database-1.cdd6uaafebcj.ap-northeast-1.rds.amazonaws.com:5432/nasb1995?sslmode=require"
	EntClient, err = ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// オートマチック遷移する
	if err := EntClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
