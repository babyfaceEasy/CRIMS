package routes

import (
	"fmt"
	"net/http"

	"github.com/babyfaceeasy/crims/internal/handlers"
	"github.com/babyfaceeasy/crims/internal/middlewares"
	"github.com/babyfaceeasy/crims/internal/services"
	"github.com/gin-gonic/gin"
)

var (
	h handlers.Handler
	m middlewares.Middleware
)

func RegisterRoutes(engine *gin.Engine, svc services.ServiceInterface) {
	// TODO: add cors middleware later
	// engine.Use()

	h = handlers.NewHandler(svc)
	m = middlewares.NewMiddleware()

	v1 := engine.Group("/v1")

	registerCustomerRoutes(v1)
	registerCloudResourceRoutes(v1)

	engine.GET("/", func(ctx *gin.Context) {
		R := handlers.ResponseFormat{}

		app_name := "CIMS"
		R.Message = fmt.Sprint("Welcome to ", app_name)
		ctx.JSON(h.Response(http.StatusOK, R))
	})

	engine.NoRoute(m.Throttle(4), func(ctx *gin.Context) {
		R := handlers.ResponseFormat{}
		R.Message = "Page not found"
		ctx.JSON(h.Response(http.StatusNotFound, R))
	})
}
