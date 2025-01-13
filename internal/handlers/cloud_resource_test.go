package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/babyfaceeasy/crims/internal/models"
	"github.com/babyfaceeasy/crims/internal/validators"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddCloudResourcesToCustomerHandler(t *testing.T) {
	mockSvc := new(MockService)
	handler := NewHandler(mockSvc)

	t.Run("should return 400 if id is missing in the URL", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/v1/customers//cloud-resources", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/v1/customers/:id/cloud-resources", handler.AddCloudResourcesToCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 400 if request body is invalid", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/v1/customers/test-uid/cloud-resources", bytes.NewBuffer([]byte(`invalid`)))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/v1/customers/:id/cloud-resources", handler.AddCloudResourcesToCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 400 if customer does not exist", func(t *testing.T) {
		mockSvc.On("GetCustomerByUID", "test-uid").Return((*models.Customer)(nil), fmt.Errorf("customer not found")).Once()

		payload := validators.AddCloudResourcesInput{Resources: []string{"Resource1", "Resource2"}}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/v1/customers/test-uid/cloud-resources", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/v1/customers/:id/cloud-resources", handler.AddCloudResourcesToCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 400 if adding cloud resources fails", func(t *testing.T) {
		customer := models.Customer{ID: 1, Name: "John Doe", Email: "john@example.com"}
		mockSvc.On("GetCustomerByUID", "test-uid").Return(&customer, nil).Once()
		mockSvc.On("AddCloudResourcesToCustomer", customer.ID, []string{"Resource1", "Resource2"}).Return(fmt.Errorf("failed to add resources")).Once()

		payload := validators.AddCloudResourcesInput{Resources: []string{"Resource1", "Resource2"}}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/v1/customers/test-uid/cloud-resources", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/v1/customers/:id/cloud-resources", handler.AddCloudResourcesToCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should successfully add cloud resources to customer", func(t *testing.T) {
		customer := models.Customer{ID: 1, Name: "Abba Khan", Email: "khan@example.com"}
		mockSvc.On("GetCustomerByUID", "test-uid").Return(&customer, nil).Once()
		mockSvc.On("AddCloudResourcesToCustomer", customer.ID, []string{"Resource1", "Resource2"}).Return(nil).Once()

		payload := validators.AddCloudResourcesInput{Resources: []string{"Resource1", "Resource2"}}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPost, "/v1/customers/test-uid/cloud-resources", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.POST("/v1/customers/:id/cloud-resources", handler.AddCloudResourcesToCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
	})
}

func TestFetchCloudResourcesForCustomerHandler(t *testing.T) {
	mockSvc := new(MockService)
	handler := NewHandler(mockSvc)

	t.Run("should return 400 if id is missing in the URL", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/v1/customers//cloud-resources", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/v1/customers/:id/cloud-resources", handler.FetchCloudResourcesForCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 400 if customer does not exist", func(t *testing.T) {
		mockSvc.On("GetCustomerByUID", "test-uid").Return((*models.Customer)(nil), fmt.Errorf("customer not found")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/v1/customers/test-uid/cloud-resources", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/v1/customers/:id/cloud-resources", handler.FetchCloudResourcesForCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should successfully fetch cloud resources for a customer", func(t *testing.T) {
		customer := &models.Customer{
			UID:   "test-uid",
			Name:  "John Doe",
			Email: "john@example.com",
			CloudResources: []*models.CloudResource{
				{UID: "cloud123", Name: "Resource A", Type: "Type1", Region: "US-East"},
				{UID: "cloud124", Name: "Resource B", Type: "Type2", Region: "EU-West"},
			},
		}

		mockSvc.On("GetCustomerByUID", "test-uid").Return(customer, nil).Once()

		req, _ := http.NewRequest(http.MethodGet, "/v1/customers/test-uid/cloud-resources", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/v1/customers/:id/cloud-resources", handler.FetchCloudResourcesForCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response map[string]interface{}
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatal("Failed to parse response", err)
		}

		// Ensure the message is correct
		assert.Equal(t, "customer cloud resources", response["message"])

		// Ensure that the cloud resources are returned correctly
		data, ok := response["data"].([]interface{})
		assert.True(t, ok)
		assert.Len(t, data, 2)
		assert.Equal(t, "Resource A", data[0].(map[string]interface{})["name"])
		assert.Equal(t, "Resource B", data[1].(map[string]interface{})["name"])
	})

	t.Run("should return 400 if id is invalid", func(t *testing.T) {
		mockSvc.On("GetCustomerByUID", "invalid-uid").Return((*models.Customer)(nil), fmt.Errorf("customer not found")).Once()

		req, _ := http.NewRequest(http.MethodGet, "/v1/customers/invalid-uid/cloud-resources", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.GET("/v1/customers/:id/cloud-resources", handler.FetchCloudResourcesForCustomer)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestUpdateCloudResourceHandler(t *testing.T) {
	mockSvc := new(MockService)
	handler := NewHandler(mockSvc)

	t.Run("should return 400 if cloud resource id is missing in the URL", func(t *testing.T) {
		mockSvc.On("GetCustomerByUID", "invalid-uid").Return((*models.Customer)(nil), fmt.Errorf("customer not found")).Once()

		req, _ := http.NewRequest(http.MethodPut, "/v1/cloud-resources//update", nil)
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/v1/cloud-resources/:id/update", handler.UpdateCloudResource)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 400 if request body is invalid", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPut, "/v1/cloud-resources/test-uid/update", bytes.NewBuffer([]byte(`invalid`)))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/v1/cloud-resources/:id/update", handler.UpdateCloudResource)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 500 if cloud resource fetch fails", func(t *testing.T) {
		mockSvc.On("GetCloudResourceByUID", "test-uid").Return(nil, fmt.Errorf("cloud resource not found")).Once()

		payload := validators.UpdateCloudResourceInput{Name: "New Resource", Type: "Type1", Region: "US-East"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/cloud-resources/test-uid/update", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/v1/cloud-resources/:id/update", handler.UpdateCloudResource)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should return 400 if the cloud resource name is already in use", func(t *testing.T) {
		existingResource := &models.CloudResource{UID: "test-uid", Name: "Existing Resource", Type: "Type1", Region: "US-East"}
		mockSvc.On("GetCloudResourceByUID", "test-uid").Return(existingResource, nil).Once()
		mockSvc.On("IsCloudResourceNameAvailable", "New Resource").Return(false, nil).Once()

		payload := validators.UpdateCloudResourceInput{Name: "New Resource", Type: "Type1", Region: "US-East"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/cloud-resources/test-uid/update", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/v1/cloud-resources/:id/update", handler.UpdateCloudResource)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return 500 if updating cloud resource fails", func(t *testing.T) {
		existingResource := &models.CloudResource{UID: "test-uid", Name: "Existing Resource", Type: "Type1", Region: "US-East"}
		mockSvc.On("GetCloudResourceByUID", "test-uid").Return(existingResource, nil).Once()
		mockSvc.On("IsCloudResourceNameAvailable", "New Resource").Return(true, nil).Once()
		mockSvc.On("UpdateCloudResource", existingResource, existingResource.ID).Return(fmt.Errorf("failed to update resource")).Once()

		payload := validators.UpdateCloudResourceInput{Name: "New Resource", Type: "Type1", Region: "US-East"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/cloud-resources/test-uid/update", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/v1/cloud-resources/:id/update", handler.UpdateCloudResource)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
	})

	t.Run("should successfully update the cloud resource", func(t *testing.T) {
		existingResource := &models.CloudResource{UID: "test-uid", Name: "Existing Resource", Type: "Type1", Region: "US-East"}
		mockSvc.On("GetCloudResourceByUID", "test-uid").Return(existingResource, nil).Once()
		mockSvc.On("IsCloudResourceNameAvailable", "New Resource").Return(true, nil).Once()
		mockSvc.On("UpdateCloudResource", existingResource, existingResource.ID).Return(nil).Once()

		payload := validators.UpdateCloudResourceInput{Name: "New Resource", Type: "Type1", Region: "US-East"}
		body, _ := json.Marshal(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/cloud-resources/test-uid/update", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()
		router := gin.Default()
		router.PUT("/v1/cloud-resources/:id/update", handler.UpdateCloudResource)
		router.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response map[string]interface{}
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatal("Failed to parse response", err)
		}

		// Ensure the message is correct
		assert.Equal(t, "cloud resource update was successful", response["message"])

		// Ensure the updated resource data is returned
		data := response["data"].(map[string]interface{})
		assert.Equal(t, "New Resource", data["name"])
		assert.Equal(t, "Type1", data["type"])
		assert.Equal(t, "US-East", data["region"])
	})
}
