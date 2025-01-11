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
	DoesCloudResourceExist(resource string) (bool, error)
	GetCloudResourcesByCustomerID(customerID uint) ([]models.CloudResource, error)
	UpdateCloudResource(resource models.CloudResource, cloudResourceID uint) error
	DeleteCloudResource(cloudResourceID uint) error
}
