package repository

import (
	"testing"

	"github.com/babyfaceeasy/crims/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Customer{}, &models.CloudResource{})
	return db
}

func TestSaveCustomer(t *testing.T) {
	db := setupTestDB()
	repo := Repository{DB: db}

	t.Run("should successfully save a customer", func(t *testing.T) {
		err := repo.SaveCustomer("John Doe", "john@example.com")
		assert.NoError(t, err)

		var customer models.Customer
		err = db.First(&customer, "email = ?", "john@example.com").Error
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", customer.Name)
		assert.Equal(t, "john@example.com", customer.Email)
	})
}

func TestIsEmailTaken(t *testing.T) {
	db := setupTestDB()
	repo := Repository{DB: db}

	// Seed data
	db.Create(&models.Customer{Name: "John Doe", Email: "john@example.com"})

	t.Run("should return true if email is taken", func(t *testing.T) {
		taken, err := repo.IsEmailTaken("john@example.com")
		assert.NoError(t, err)
		assert.True(t, taken)
	})

	t.Run("should return false if email is not taken", func(t *testing.T) {
		taken, err := repo.IsEmailTaken("jane@example.com")
		assert.NoError(t, err)
		assert.False(t, taken)
	})
}

func TestGetCustomer(t *testing.T) {
	db := setupTestDB()
	repo := Repository{DB: db}

	// Seed data
	customer := models.Customer{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	db.Create(&customer)

	t.Run("should successfully get a customer by email", func(t *testing.T) {
		result, err := repo.GetCustomer(nil, "email = ?", "john@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "john@example.com", result.Email)
	})

	t.Run("should return error if customer does not exist", func(t *testing.T) {
		result, err := repo.GetCustomer(nil, "email = ?", "jane@example.com")
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}
