package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl   string
	AppPort string
}

func Load() Config {
	//Load .env
	_ = godotenv.Load()

	cfg := Config{
		DBUrl:   os.Getenv("DATABASE_URL"),
		AppPort: os.Getenv("APP_PORT"),
	}

	if cfg.DBUrl == "" {
		log.Fatal("DATABASE_URL is not set in .env")
	}
	if cfg.AppPort == "" {
		cfg.AppPort = "8080" //Default fallback
	}

	return cfg
}
