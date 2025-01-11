package handlers

import (
	"net/http"

	"github.com/babyfaceeasy/crims/internal/messages"
	"github.com/babyfaceeasy/crims/internal/validators"
	"github.com/gin-gonic/gin"
)

func (h Handler) AddCloudResourcesToCustomer(ctx *gin.Context) {
	R := ResponseFormat{}

	//customer_id := ctx.Param("id")

	var i validators.AddResourcesInput
	if err := ctx.ShouldBindJSON(&i); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	// check if resources exists
	//h.svc.AddCustomer()

	R.Message = "add cloud resources to customer"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) FetchCloudResourcesForCustomer(ctx *gin.Context) {
	R := ResponseFormat{}

	R.Message = "customer cloud resources"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) UpdateCloudResource(ctx *gin.Context) {
	R := ResponseFormat{}

	R.Message = "update cloud resources"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) DeleteCloudResource(ctx *gin.Context) {
	R := ResponseFormat{}

	R.Message = "delete cloud resources"
	ctx.JSON(h.Response(http.StatusOK, R))
}
