package routes

import "github.com/gin-gonic/gin"

func registerCloudResourceRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/cloud-resources")
	router.Use()
	{
		router.PUT("/:id", m.Throttle(2), h.UpdateCloudResource)
		router.DELETE("/:id", m.Throttle(2), h.DeleteCloudResource)
	}
}
