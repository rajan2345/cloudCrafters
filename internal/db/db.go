package db

import (
	"log"

	"cloudCrafters/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v : ", err)
	}
	log.Println("Connected to Postgres successfully")
	return db
}
