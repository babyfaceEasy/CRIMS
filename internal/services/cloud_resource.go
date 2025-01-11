package services

import (
	"github.com/babyfaceeasy/crims/internal/models"
)

func (svc Service) DoesCloudResourceExist(resource string) (bool, error) {
	return svc.repo.DoesCloudResourceExist(resource)
}

func (svc Service) GetCloudResourcesByCustomerID(customerID uint) ([]models.CloudResource, error) {
	return svc.repo.GetCloudResourcesByCustomerID(customerID)
}

func (svc Service) UpdateCloudResource(resource models.CloudResource, cloudResourceID uint) error {
	return svc.repo.UpdateCloudResource(resource, cloudResourceID)
}

func (svc Service) DeleteCloudResource(cloudResourceID uint) error {
	return svc.repo.DeleteCloudResource(cloudResourceID)
}
