package seeds

import (
	"log"

	"github.com/babyfaceeasy/crims/internal/data"
	"github.com/babyfaceeasy/crims/internal/models"
)

func (s Seed) CustomerSeeder() {
	for _, customer := range data.Customers {
		err := s.repo.DB.Table("customers").Where(models.Customer{
			Email: customer.Email,
		}).Assign(models.Customer{
			Email: customer.Email,
		}).FirstOrCreate(&models.Customer{
			Email: customer.Email,
		}).Error
		if err != nil {
			log.Fatalf("error in seeding customers table y: %v", err)
		}
	}
}
