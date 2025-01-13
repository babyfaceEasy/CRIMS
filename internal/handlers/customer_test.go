package handlers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomerHandler(t *testing.T) {
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
	handler := NewHandler(mockSvc)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/v1/customers", handler.CreateCustomer)

	t.Run("should return bad request for invalid payload", func(t *testing.T) {
		// Invalid JSON payload
		payload := `{"name": "John Doe", "email": "invalid-email"}`
		req, _ := http.NewRequest(http.MethodPost, "/v1/customers", bytes.NewBufferString(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return internal server error if checking email fails", func(t *testing.T) {
		payload := `{"name": "John Doe", "email": "john@example.com"}`
		mockSvc.On("IsEmailTaken", "john@example.com").Return(false, errors.New("database error")).Once()

		req, _ := http.NewRequest(http.MethodPost, "/v1/customers", bytes.NewBufferString(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockSvc.AssertCalled(t, "IsEmailTaken", "john@example.com")
	})

	t.Run("should return unprocessable entity if email is already taken", func(t *testing.T) {
		payload := `{"name": "John Doe", "email": "john@example.com"}`
		mockSvc.On("IsEmailTaken", "john@example.com").Return(true, nil).Once()

		req, _ := http.NewRequest(http.MethodPost, "/v1/customers", bytes.NewBufferString(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		mockSvc.AssertCalled(t, "IsEmailTaken", "john@example.com")
	})

	t.Run("should return internal server error if adding customer fails", func(t *testing.T) {
		payload := `{"name": "John Doe", "email": "john@example.com"}`
		mockSvc.On("IsEmailTaken", "john@example.com").Return(false, nil).Once()
		mockSvc.On("AddCustomer", "John Doe", "john@example.com").Return(errors.New("failed to insert customer")).Once()

		req, _ := http.NewRequest(http.MethodPost, "/v1/customers", bytes.NewBufferString(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockSvc.AssertCalled(t, "IsEmailTaken", "john@example.com")
		mockSvc.AssertCalled(t, "AddCustomer", "John Doe", "john@example.com")
	})

	t.Run("should create customer successfully", func(t *testing.T) {
		payload := `{"name": "John Doe", "email": "john@example.com"}`
		mockSvc.On("IsEmailTaken", "john@example.com").Return(false, nil).Once()
		mockSvc.On("AddCustomer", "John Doe", "john@example.com").Return(nil).Once()

		req, _ := http.NewRequest(http.MethodPost, "/v1/customers", bytes.NewBufferString(payload))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockSvc.AssertCalled(t, "IsEmailTaken", "john@example.com")
		mockSvc.AssertCalled(t, "AddCustomer", "John Doe", "john@example.com")
	})
}
