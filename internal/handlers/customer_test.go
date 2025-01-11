package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/babyfaceeasy/crims/internal/validators"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCustomersHandler(t *testing.T) {
	svc := new(MockService)
	handler := NewHandler(svc)

	t.Run("should return 200 status code when called", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/v1/customers", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.GET("/v1/customers", handler.GetCustomers)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestCreateCustomerHandler(t *testing.T) {
	mockSvc := new(MockService)
	mockSvc.On("IsEmailTaken", "john@example.com").Return(false, nil).Once()
	mockSvc.On("AddCustomer", "John Doe", "john@example.com").Return(nil).Once()
	handler := NewHandler(mockSvc)

	t.Run("should correctly create the customer", func(t *testing.T) {
		payload := validators.CreateCustomerInput{
			Name:  "John Doe",
			Email: "john@example.com",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/v1/customers", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/v1/customers", handler.CreateCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := validators.CreateCustomerInput{
			Name:  "John Doe",
			Email: "johnexample.com",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/v1/customers", handler.CreateCustomer)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

// MockService is a mock implementation of ServiceInterface
type MockService struct {
	mock.Mock
}

// AddCustomer mocks the AddCustomer method
func (m *MockService) AddCustomer(name, email string) error {
	args := m.Called(name, email)
	return args.Error(0)
}

// IsEmailTaken mocks the IsEmailTaken method
func (m *MockService) IsEmailTaken(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}
