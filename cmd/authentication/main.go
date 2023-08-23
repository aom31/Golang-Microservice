package main

import (
	"context"
	"fmt"

	"golang-microservice/config"
	"golang-microservice/db"
)

func main() {

	cfg := config.NewConfig()
	//connect db
	dbClientMongo := db.DBConn(cfg)
	defer dbClientMongo.Disconnect(context.Background())

	fmt.Println(cfg)
}
