package services

import (
	"errors"

	"github.com/babyfaceeasy/crims/internal/models"
	"gorm.io/gorm"
)

func (svc Service) DoesCloudResourceExist(resource string) (bool, error) {
	return svc.repo.DoesCloudResourceExist(resource)
}

func (svc Service) GetCloudResourcesByCustomerID(customerID uint) ([]models.CloudResource, error) {
	cloudResources, err := svc.repo.GetCloudResources(nil, "customer_id = ?", customerID)
	if err != nil {
		return nil, err
	}
	return cloudResources, nil
}

func (svc Service) GetCloudResourceByUID(cloudResourceUID string) (models.CloudResource, error) {
	cr, err := svc.repo.GetCloudResource(nil, "uid = ?", cloudResourceUID)
	if err != nil {
		return models.CloudResource{}, err
	}
	return *cr, nil
}

func (svc Service) GetCloudResourceByName(cloudResourceName string) (*models.CloudResource, error) {
	cr, err := svc.repo.GetCloudResource(nil, "name = ?", cloudResourceName)
	if err != nil {
		return nil, err
	}
	return cr, nil
}

func (svc Service) IsCloudResourceNameAvailable(cloudResourceName string) (bool, error) {
	cr, err := svc.repo.GetCloudResource(nil, "name = ?", cloudResourceName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return cr.Name == cloudResourceName, nil
}

func (svc Service) UpdateCloudResource(resource models.CloudResource, cloudResourceID uint) error {
	return svc.repo.UpdateCloudResource(nil, resource, cloudResourceID)
}

func (svc Service) DeleteCloudResource(cloudResourceID uint) error {
	return svc.repo.DeleteCloudResource(nil, cloudResourceID)
}

func (svc Service) AddCloudResourcesToCustomer(customerID uint, resourceNames []string) error {
	return svc.repo.AttachCloudResourcesByNames(nil, customerID, resourceNames)
}
