package service

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"log"
	"newdeal/ent"
	"time"
)

var EntCore *ent.Client
var err error

func InitEntClient() {
	// DSNを定義する
	dsn := "postgres://postgres:postgres@52.198.76.7:5432/nasb?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// 🎯 ここでプール設定
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(33 * time.Minute)
	// EntClientを取得する
	drv := entsql.OpenDB(dialect.Postgres, db)
	EntCore := ent.NewClient(ent.Driver(drv))
	// オートマチック遷移する
	if err := EntCore.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
