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
	dsn := "postgres://postgres:postgres@52.198.76.7:5432/nasb?sslmode=disable"
	EntClient, err = ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// オートマチック遷移する
	if err := EntClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
