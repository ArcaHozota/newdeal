package service

import (
	"context"
	"log"
	"newdeal/ent"
)

var EntClient *ent.Client
var err error

func init() {
	// entを配置する
	EntClient, err = ent.Open("postgres", "host=<52.199.163.241> port=<5432> user=<postgres> dbname=<nasb1995> password=<postgres>")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer EntClient.Close()
	// Run the auto migration tool.
	if err := EntClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
