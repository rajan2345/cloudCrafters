package mappings

import "gorm.io/gorm"

type MappingRepository struct {
	DB *gorm.DB
}

func NewMappingRepository(db *gorm.DB) *MappingRepository {
	return &MappingRepository{DB: db}
}

// GetMapping returns ONE Mapping based on:
// from_provider, to_provider, from_code
func (r *MappingRepository) GetMapping(fromProvider, fromCode, toProvider string) (*Mappings, error) {
	var mapping Mappings
	result := r.DB.Where("from_provider=? AND from_code = ? AND to_provider = ?", fromProvider, fromCode, toProvider).Find(&mapping)
	if result.Error != nil {
		return nil, result.Error
	}
	return &mapping, nil
}
