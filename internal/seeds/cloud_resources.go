package seeds

import (
	"log"

	"github.com/babyfaceeasy/crims/internal/data"
	"github.com/babyfaceeasy/crims/internal/models"
)

func (s Seed) CloudResourceSeeder() {
	customer2 := models.Customer{Name: "Jane Smith", Email: "jane.smith@example.com"}
	err := s.repo.DB.Table("customers").Where(customer2).Assign(customer2).FirstOrCreate(&customer2).Error
	if err != nil {
		log.Fatalf("error in seeding customers table for CloudResources y: %v", err)
	}

	for _, resource := range data.CloudResources {
		s.repo.DB.Create(&resource)
		s.repo.DB.Model(&customer2).Association("CloudResources").Append(&resource)
	}
}
