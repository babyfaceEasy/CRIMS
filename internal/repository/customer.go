package repository

import (
	"fmt"

	"github.com/babyfaceeasy/crims/internal/models"
	"gorm.io/gorm"
)

func (repo Repository) SaveCustomer(name, email string) error {
	return repo.DB.Create(&models.Customer{
		Name:  name,
		Email: email,
	}).Error
}

func (repo Repository) IsEmailTaken(email string) (bool, error) {
	var c models.Customer
	result := repo.DB.Where("email = ?", email).First(&c)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, fmt.Errorf("error in getting / fetching customer: %v", result.Error)
	}
	return true, nil
}
