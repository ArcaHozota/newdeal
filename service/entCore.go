package service

import (
	"context"
	_ "github.com/lib/pq"
	"log"
	"newdeal/ent"
)

var EntCore *ent.Client
var err error

func InitEntClient() {
	// DSNを定義する
	dsn := "postgres://postgres:postgres@52.198.76.7:5432/nasb?sslmode=disable"
	EntCore, err = ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// オートマチック遷移する
	if err := EntCore.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
