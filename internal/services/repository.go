package services

import (
	"gorm.io/gorm"
)

type ServiceRepository struct {
	DB *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{DB: db}
}

func (r *ServiceRepository) GetAll() ([]Services, error) {
	var services []Services
	result := r.DB.Find(&services)
	return services, result.Error
}

func (r *ServiceRepository) GetByProvider(provider string) ([]Services, error) {
	var services []Services
	result := r.DB.Where("provider = ?", provider).Find(&services)
	return services, result.Error
}

func (r *ServiceRepository) Get(provider, code string) (*Services, error) {
	var services Services
	result := r.DB.Where("provider = ? AND code = ?", provider, code).First(&services)

	if result.Error != nil {
		return nil, result.Error
	}

	return &services, nil
}
