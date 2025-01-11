package repository

import (
	"errors"
	"fmt"

	"github.com/babyfaceeasy/crims/internal/models"
	"gorm.io/gorm"
)

func (repo Repository) AttachCloudResourcesByNames(db *gorm.DB, customerID uint, resourceNames []string) error {
	if db == nil {
		db = repo.DB
	}
	var customer models.Customer
	if err := db.First(&customer, customerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("customer with ID %d not found", customerID)
		}
		return fmt.Errorf("failed to retrieve customer: %w", err)
	}

	// Check for missing resource names
	var resources []models.CloudResource
	if err := db.Where("name IN ?", resourceNames).Find(&resources).Error; err != nil {
		return fmt.Errorf("failed to retrieve cloud resources: %w", err)
	}

	// Check for missing resource names
	foundNames := make(map[string]bool)
	for _, resource := range resources {
		foundNames[resource.Name] = true
	}
	var missingNames []string
	for _, name := range resourceNames {
		if !foundNames[name] {
			missingNames = append(missingNames, name)
		}
	}
	if len(missingNames) > 0 {
		return fmt.Errorf("the following cloud resources do not exist: %v", missingNames)
	}

	// Attach the resources to the customer
	if err := db.Model(&customer).Association("CloudResources").Append(resources); err != nil {
		return fmt.Errorf("failed to attach cloud resources to customer: %w", err)
	}

	return nil
}

func (repo Repository) DoesCloudResourceExist(resourceName string) (bool, error) {
	var c models.CloudResource
	result := repo.DB.Where("name = ?", resourceName).First(&c)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("error in getting / fetching cloud resource: %v", result.Error)
	}
	return true, nil
}

func (repo Repository) GetCloudResourcesByCustomerID(customerID uint) ([]models.CloudResource, error) {
	resources := []models.CloudResource{}

	err := repo.DB.Model(models.CloudResource{}).Where("customer_id = ?", customerID).Find(&resources).Error
	if err != nil {
		return resources, fmt.Errorf("error in getting cloud resources: %v", err)
	}

	return resources, nil
}

func (repo Repository) UpdateCloudResource(tx *gorm.DB, cloudResource models.CloudResource, cloudResourceID uint) error {
	if tx == nil {
		tx = repo.DB
	}

	return tx.Where("id = ?", cloudResource.ID).Updates(&cloudResource).Error
}

func (repo Repository) DeleteCloudResourceOLD(tx *gorm.DB, cloudResource models.CloudResource) error {
	if tx == nil {
		tx = repo.DB
	}

	if err := tx.Delete(cloudResource).Error; err != nil {
		return err
	}

	return nil
}

func (repo Repository) DeleteCloudResource(tx *gorm.DB, resourceID uint) error {
	if tx == nil {
		tx = repo.DB
	}
	var resource models.CloudResource
	if err := tx.First(&resource, resourceID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cloud resource with ID %d not found", resourceID)
		}
		return err
	}

	// Delete the cloud resource
	if err := tx.Delete(&resource).Error; err != nil {
		return fmt.Errorf("failed to delete cloud resource: %w", err)
	}

	return nil
}

func (repo Repository) GetCloudResource(tx *gorm.DB, query interface{}, args ...interface{}) (*models.CloudResource, error) {
	if tx == nil {
		tx = repo.DB
	}
	cr := models.CloudResource{}
	err := tx.Model(models.CloudResource{}).Preload("Customers").Where(query, args...).First(&cr).Error
	if err != nil {
		return nil, err
	}

	return &cr, nil
}

func (repo Repository) GetCloudResources(tx *gorm.DB, query interface{}, args ...interface{}) ([]models.CloudResource, error) {
	crs := []models.CloudResource{}
	if tx == nil {
		tx = repo.DB
	}

	err := tx.Model(models.CloudResource{}).
		Preload("Customers").
		Where(query, args...).Find(&crs).Error
	if err != nil {
		return nil, fmt.Errorf("error in getting cloud resources %w", err)
	}

	return crs, nil
}
