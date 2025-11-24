package db

import (
	"cloudCrafters/internal/mappings"
	"cloudCrafters/internal/services"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&services.Services{})
	db.AutoMigrate(&mappings.Mappings{})
}
