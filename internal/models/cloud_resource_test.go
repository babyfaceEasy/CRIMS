package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCloudResourceModel(t *testing.T) {
	db := setupTestDB()

	t.Run("should create a cloud resource", func(t *testing.T) {
		resource := CloudResource{
			Name:   "TestResource",
			Type:   "Compute",
			Region: "us-west-1",
		}

		err := db.Create(&resource).Error
		assert.NoError(t, err)
		assert.NotZero(t, resource.ID, "CloudResource ID should be set after creation")
	})

	t.Run("should enforce unique name constraint", func(t *testing.T) {
		resource1 := CloudResource{
			Name:   "UniqueResource",
			Type:   "Storage",
			Region: "us-east-1",
		}
		resource2 := CloudResource{
			Name:   "UniqueResource",
			Type:   "Compute",
			Region: "us-west-2",
		}

		err := db.Create(&resource1).Error
		assert.NoError(t, err)

		err = db.Create(&resource2).Error
		assert.Error(t, err, "should fail due to unique name constraint")
	})

	t.Run("should test many-to-many relationship with customers", func(t *testing.T) {
		customer := Customer{
			Name:  "John Doe",
			Email: "john@example.com",
		}
		resource1 := CloudResource{Name: "Resource1", Type: "Compute", Region: "us-west-1"}
		resource2 := CloudResource{Name: "Resource2", Type: "Storage", Region: "us-east-1"}

		// Create a customer and resources
		err := db.Create(&customer).Error
		assert.NoError(t, err)

		err = db.Create(&resource1).Error
		assert.NoError(t, err)

		err = db.Create(&resource2).Error
		assert.NoError(t, err)

		// Associate resources with the customer
		err = db.Model(&customer).Association("CloudResources").Append(&resource1, &resource2)
		assert.NoError(t, err)

		// Retrieve resource with associated customers
		var retrievedResource CloudResource
		err = db.Preload("Customers").First(&retrievedResource, resource1.Name).Error
		assert.NoError(t, err)
		assert.Len(t, retrievedResource.Customers, 1, "Resource1 should have 1 associated customer")
		assert.Equal(t, customer.Name, retrievedResource.Customers[0].Name, "Associated customer name should match")
	})

	t.Run("should update a cloud resource", func(t *testing.T) {
		resource := CloudResource{
			Name:   "OldResource",
			Type:   "Compute",
			Region: "us-west-1",
		}
		err := db.Create(&resource).Error
		assert.NoError(t, err)

		// Update the resource
		resource.Name = "UpdatedResource"
		err = db.Save(&resource).Error
		assert.NoError(t, err)

		// Verify the update
		var updatedResource CloudResource
		err = db.First(&updatedResource, resource.Name).Error
		assert.NoError(t, err)
		assert.Equal(t, "UpdatedResource", updatedResource.Name, "Resource name should be updated")
	})

	t.Run("should delete a cloud resource", func(t *testing.T) {
		resource := CloudResource{
			Name:   "DeletableResource",
			Type:   "Storage",
			Region: "us-east-1",
		}

		// Create the resource
		err := db.Create(&resource).Error
		assert.NoError(t, err)

		// Ensure the resource exists in the database
		var found CloudResource
		err = db.First(&found, resource.Name).Error
		assert.NoError(t, err, "Resource should exist in the database before deletion")
		assert.Equal(t, resource.Name, found.Name, "Resource name should match before deletion")

		// Delete the resource
		err = db.Delete(&resource).Error
		assert.NoError(t, err, "Resource deletion should succeed")

		// Confirm the resource is deleted by attempting to retrieve it
		err = db.First(&found, resource.Name).Error
		assert.Error(t, err, "Resource should not be found in the database after deletion")
		assert.Equal(t, gorm.ErrRecordNotFound, err, "Expected error to be ErrRecordNotFound after deletion")
	})
}
