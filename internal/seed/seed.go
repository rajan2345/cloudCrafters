package seed

import (
	"errors"
	"log"

	"cloudCrafters/internal/mappings"
	"cloudCrafters/internal/services"

	"gorm.io/gorm"
)

// Run seeds initial services and mappings into the database
func Run(db *gorm.DB) {
	seedServices(db)
	seedMappings(db)
}

func seedServices(db *gorm.DB) {
	servicesData := []services.Services{

		//AWS
		{Provider: "aws", Code: "ec2", Name: "Amazon EC2", Category: "compute"},
		{Provider: "aws", Code: "s3", Name: "Amazon S3", Category: "storage"},
		{Provider: "aws", Code: "rds-mysql", Name: "Amazon RDS MySQL", Category: "database"},

		//Azure
		{Provider: "azure", Code: "virtual-machines", Name: "Azure Virtual Machines", Category: "compute"},
		{Provider: "azure", Code: "blob-storage", Name: "Azure Blob Storage ", Category: "storage"},
		{Provider: "azure", Code: "azure-mysql", Name: "Azure Database for MySQL", Category: "database"},

		//GCP
		{Provider: "gcp", Code: "compute-engine", Name: "Compute Engine", Category: "compute"},
		{Provider: "gcp", Code: "cloud-storage", Name: "Cloud Storage", Category: "storage"},
		{Provider: "gcp", Code: "cloud-sql-mysql", Name: "Cloud SQL MySQL", Category: "database"},
	}

	for _, svc := range servicesData {
		var existing services.Services
		result := db.Where("provider = ? AND code = ?", svc.Provider, svc.Code).First(&existing)

		// if not found → insert
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.Create(&svc)
		}
	}
	log.Println("Seeded services succe")
}

func seedMappings(db *gorm.DB) {
	mappingsData := []mappings.Mappings{
		// EC2 mappings
		{FromProvider: "aws", FromCode: "ec2", ToProvider: "azure", ToCode: "virtual-machines"},
		{FromProvider: "aws", FromCode: "ec2", ToProvider: "gcp", ToCode: "compute-engine"},

		// S3 mappings
		{FromProvider: "aws", FromCode: "s3", ToProvider: "azure", ToCode: "blob-storage"},
		{FromProvider: "aws", FromCode: "s3", ToProvider: "gcp", ToCode: "cloud-storage"},

		//rds mapping
		// RDS MySQL → Managed Database MySQL
		{FromProvider: "aws", FromCode: "rds-mysql", ToProvider: "azure", ToCode: "azure-mysql"},
		{FromProvider: "aws", FromCode: "rds-mysql", ToProvider: "gcp", ToCode: "cloud-sql-mysql"},
	}

	for _, mp := range mappingsData {
		var existing mappings.Mappings
		err := db.Where(
			"from_provider = ? AND from_code = ? AND to_provider = ?",
			mp.FromProvider, mp.FromCode, mp.ToProvider,
		).First(&existing).Error

		if err == gorm.ErrRecordNotFound {
			db.Create(&mp)
		}
	}
	log.Println("Seeded mappings successfully")
}
