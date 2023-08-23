package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB *Db
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DB: &Db{
			DBName: os.Getenv("DATABASE_NAME"),
			DBUrl:  os.Getenv("DB_URL"),
			DBPort: os.Getenv("DATABASE_PORT"),
		},
	}
}

type Db struct {
	DBName string
	DBUrl  string
	DBPort string
}
