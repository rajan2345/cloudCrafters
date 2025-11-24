package services

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	Provider string `json:"provider"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
