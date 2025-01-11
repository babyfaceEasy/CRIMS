package services

import "github.com/babyfaceeasy/crims/internal/models"

func (svc Service) AddCustomer(name, email string) error {
	return svc.repo.SaveCustomer(name, email)
}

func (svc Service) IsEmailTaken(email string) (bool, error) {
	return svc.repo.IsEmailTaken(email)
}

func (svc Service) GetCustomerByUID(customerUID string) (*models.Customer, error) {
	customer, err := svc.repo.GetCustomer(nil, "uid = ?", customerUID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
