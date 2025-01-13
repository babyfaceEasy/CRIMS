package handlers

import (
	"github.com/babyfaceeasy/crims/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockService is a mock implementation of ServiceInterface
type MockService struct {
	mock.Mock
}

func (m *MockService) AddCustomer(name, email string) error {
	args := m.Called(name, email)
	return args.Error(0)
}

func (m *MockService) IsEmailTaken(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

func (m *MockService) GetCustomerByUID(customerUID string) (*models.Customer, error) {
	args := m.Called(customerUID)
	return args.Get(0).(*models.Customer), args.Error(1)
}

func (m *MockService) IsCloudResourceNameAvailable(cloudResourceName string) (bool, error) {
	args := m.Called(cloudResourceName)
	return args.Bool(0), args.Error(1)
}

func (m *MockService) GetCloudResourcesByCustomerID(customerID uint) ([]models.CloudResource, error) {
	args := m.Called(customerID)
	return args.Get(0).([]models.CloudResource), args.Error(1)
}

func (m *MockService) GetCloudResourceByUID(cloudResourceUID string) (models.CloudResource, error) {
	args := m.Called(cloudResourceUID)
	return args.Get(0).(models.CloudResource), args.Error(1)
}

func (m *MockService) GetCloudResourceByName(cloudResourceName string) (*models.CloudResource, error) {
	args := m.Called(cloudResourceName)
	return args.Get(0).(*models.CloudResource), args.Error(1)
}

func (m *MockService) UpdateCloudResource(resource models.CloudResource, cloudResourceID uint) error {
	args := m.Called(resource, cloudResourceID)
	return args.Error(0)
}

func (m *MockService) DeleteCloudResource(cloudResourceID uint) error {
	args := m.Called(cloudResourceID)
	return args.Error(0)
}

func (m *MockService) AddCloudResourcesToCustomer(customerID uint, resourceNames []string) error {
	args := m.Called(customerID, resourceNames)
	return args.Error(0)
}
