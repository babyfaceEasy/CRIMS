package routes

import "github.com/gin-gonic/gin"

func registerCustomerRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/customers")
	router.Use()
	{
		router.GET("", h.GetCustomers)
		router.GET("/:id/cloud-resources", m.Throttle(4), h.FetchCloudResourcesForCustomer)
		router.POST("/:id/cloud-resources", m.Throttle(2), h.AddCloudResourcesToCustomer)
		router.POST("", m.Throttle(4), h.CreateCustomer)
	}
}
