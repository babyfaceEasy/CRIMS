package services

func (svc Service) AddCustomer(name, email string) error {
	return svc.repo.SaveCustomer(name, email)
}

func (svc Service) IsEmailTaken(email string) (bool, error) {
	return svc.repo.IsEmailTaken(email)
}
