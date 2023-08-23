package db

import (
	"golang-microservice/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
)

func DBConn(cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DB.DBUrl))
	if err != nil {
		log.Fatalf("failed connect to mongodb with %v", cfg.DB.DBUrl)
	}

	//ping check db
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("failed to ping mongodb:%s with url:%v", err.Error(), cfg.DB.DBUrl)
	}

	log.Printf("successful connected mongodb with url: %v", cfg.DB.DBUrl)

	return client
}
