package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
// Setup a test database
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the test database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Customer{}, &CloudResource{})
	if err != nil {
		panic("failed to migrate the test database")
	}

	return db
}
*/

func TestCustomerModel(t *testing.T) {
	db := setupTestDB()

	t.Run("should create a customer", func(t *testing.T) {
		customer := Customer{
			Name:  "John Doe",
			Email: "johndoe@example.com",
		}

		err := db.Create(&customer).Error
		assert.NoError(t, err)
		assert.NotZero(t, customer.ID, "Customer ID should be set after creation")
	})

	t.Run("should not allow duplicate email", func(t *testing.T) {
		customer1 := Customer{
			Name:  "Alice",
			Email: "alice@example.com",
		}
		customer2 := Customer{
			Name:  "Bob",
			Email: "alice@example.com",
		}

		err := db.Create(&customer1).Error
		assert.NoError(t, err)

		err = db.Create(&customer2).Error
		assert.Error(t, err, "should fail due to duplicate email")
	})

	t.Run("should test many-to-many relationship with CloudResource", func(t *testing.T) {
		customer := Customer{
			Name:  "Jane Doe",
			Email: "janedoe@example.com",
		}
		cloudResource1 := CloudResource{Name: "Resource1", Type: "Compute", Region: "us-west-1"}
		cloudResource2 := CloudResource{Name: "Resource2", Type: "Storage", Region: "us-east-1"}

		err := db.Create(&customer).Error
		assert.NoError(t, err)

		err = db.Create(&cloudResource1).Error
		assert.NoError(t, err)

		err = db.Create(&cloudResource2).Error
		assert.NoError(t, err)

		// Associate resources with the customer
		err = db.Model(&customer).Association("CloudResources").Append(&cloudResource1, &cloudResource2)
		assert.NoError(t, err)

		// Retrieve customer with associated resources
		var retrievedCustomer Customer
		err = db.Preload("CloudResources").First(&retrievedCustomer, customer.ID).Error
		assert.NoError(t, err)
		assert.Len(t, retrievedCustomer.CloudResources, 2, "Customer should have 2 associated resources")
	})

	t.Run("should soft delete a customer", func(t *testing.T) {
		customer := Customer{
			Name:  "Mark Doe",
			Email: "markdoe@example.com",
		}
		err := db.Create(&customer).Error
		assert.NoError(t, err)

		// Soft delete the customer
		err = db.Delete(&customer).Error
		assert.NoError(t, err)

		// Ensure the customer is not visible in queries
		var found Customer
		err = db.First(&found, customer.ID).Error
		assert.Error(t, err, "Customer should not be found after deletion")

		// Ensure the customer is still in the database (soft delete)
		err = db.Unscoped().First(&found, customer.ID).Error
		assert.NoError(t, err, "Customer should exist when using Unscoped")
		assert.NotNil(t, found.DeletedAt.Time, "DeletedAt should be set after soft delete")
	})
}
