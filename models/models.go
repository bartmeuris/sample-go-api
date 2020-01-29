package models

import "github.com/jinzhu/gorm"

// Migrate migrates all known models in the database
func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&TeamDb{}); err != nil && err.Error != nil {
		return err.Error
	}
	if err := db.AutoMigrate(&ProjectDb{}); err != nil && err.Error != nil {
		return err.Error
	}
	if err := db.AutoMigrate(&MicroserviceDb{}); err != nil && err.Error != nil {
		return err.Error
	}
	return nil
}
