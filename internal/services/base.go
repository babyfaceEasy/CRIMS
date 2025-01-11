package services

import (
	"github.com/babyfaceeasy/crims/internal/models"
	"github.com/babyfaceeasy/crims/internal/repository"
)

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) Service {
	return Service{repo: repo}
}

// ServiceInterface definition
type ServiceInterface interface {
	AddCustomer(name, email string) error
	IsEmailTaken(email string) (bool, error)
	GetCustomerByUID(customerUID string) (*models.Customer, error)
	//DoesCloudResourceExist(resource string) (bool, error)
	GetCloudResourcesByCustomerID(customerID uint) ([]models.CloudResource, error)
	GetCloudResourceByUID(cloudResourceUID string) (models.CloudResource, error)
	UpdateCloudResource(resource models.CloudResource, cloudResourceID uint) error
	DeleteCloudResource(cloudResourceID uint) error
	AddCloudResourcesToCustomer(customerID uint, resourceNames []string) error
}
