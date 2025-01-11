package services

import (
	"github.com/babyfaceeasy/crims/internal/models"
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

func (svc Service) UpdateCloudResource(resource models.CloudResource, cloudResourceID uint) error {
	cr, err := svc.repo.GetCloudResource(nil, "id = ?", cloudResourceID)
	if err != nil {
		return err
	}
	return svc.repo.UpdateCloudResource(nil, *cr, cloudResourceID)
}

func (svc Service) DeleteCloudResource(cloudResourceID uint) error {
	return svc.repo.DeleteCloudResource(nil, cloudResourceID)
}

func (svc Service) AddCloudResourcesToCustomer(customerID uint, resourceNames []string) error {
	return svc.repo.AttachCloudResourcesByNames(nil, customerID, resourceNames)
}
