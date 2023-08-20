package db

import (
	"fmt"
	"log"
	"os"
)

type Config interface {
	Dsn() string
	DbName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort string
	dbName string
	dsn    string
}

func NewConfig() Config {
	var cfg config
	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbHost = os.Getenv("DATABASE_HOST")
	cfg.dbName = os.Getenv("DATABASE_NAME")
	var err error
	cfg.dbPort = os.Getenv("DATABASE_PORT")
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}
	//mongodb: //mongoadmin:password1@0.0.0.0:3005/
	//cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%s/", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort)
	cfg.dsn = fmt.Sprintf("mongodb://mongoadmin:password1@0.0.0.0:%v/", cfg.dbPort)
	return &cfg
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DbName() string {
	return c.dbName
}
