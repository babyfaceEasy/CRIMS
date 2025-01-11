package repository

import (
	"github.com/babyfaceeasy/crims/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

type RepositoryInterface interface {
	SaveCustomer(name, email string) error
	IsEmailTaken(email string) (bool, error)
	GetCustomer(tx *gorm.DB, query interface{}, args ...interface{}) (*models.Customer, error)
	DoesCloudResourceExist(resource string) (bool, error)
	AttachCloudResourcesByNames(db *gorm.DB, customerID uint, resourceNames []string) error
	// GetCloudResourcesByCustomerID(customerID uint) ([]models.CloudResource, error)
	UpdateCloudResource(tx *gorm.DB, resource models.CloudResource, cloudResourceID uint) error
	DeleteCloudResource(tx *gorm.DB, resourceID uint) error
	GetCloudResource(tx *gorm.DB, query interface{}, args ...interface{}) (*models.CloudResource, error)
	GetCloudResources(tx *gorm.DB, query interface{}, args ...interface{}) ([]models.CloudResource, error)
}
