package migrations

import (
	"hris/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.M_Employee{},
		&models.AD_Reference{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migration completed successfully.")
}
