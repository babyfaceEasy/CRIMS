package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup a test database
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the test database")
	}

	// Migrate the schema for CloudResource and Customer
	err = db.AutoMigrate(&Customer{}, &CloudResource{})
	if err != nil {
		panic("failed to migrate the test database")
	}

	return db
}
