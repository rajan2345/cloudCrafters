package mappings

import "gorm.io/gorm"

type Mappings struct {
	gorm.Model
	FromProvider string `json:"from_provider"`
	FromCode     string `json:"from_code"`
	ToProvider   string `json:"to_provider"`
	ToCode       string `json:"to_code"`
}
