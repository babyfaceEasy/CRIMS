package handlers

import (
	"log"
	"net/http"

	"github.com/babyfaceeasy/crims/internal/messages"
	"github.com/babyfaceeasy/crims/internal/validators"
	"github.com/gin-gonic/gin"
)

func (h Handler) AddCloudResourcesToCustomer(ctx *gin.Context) {
	R := ResponseFormat{}

	customer_uid := ctx.Param("id")
	if len(customer_uid) == 0 {
		R.Error = append(R.Error, "id is required")
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	var i validators.AddCloudResourcesInput
	if err := ctx.ShouldBindJSON(&i); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	customer, err := h.svc.GetCustomerByUID(customer_uid)
	if err != nil {
		log.Printf("error fetching customer: %v", err)
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	err = h.svc.AddCloudResourcesToCustomer(customer.ID, i.Resources)
	if err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	R.Message = "add cloud resources to customer"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) FetchCloudResourcesForCustomer(ctx *gin.Context) {
	R := ResponseFormat{}

	customer_uid := ctx.Param("id")
	if len(customer_uid) == 0 {
		R.Error = append(R.Error, "id is required")
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	customer, err := h.svc.GetCustomerByUID(customer_uid)
	if err != nil {
		log.Printf("error fetching customer: %v", err)
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	R.Data = customer.CloudResources

	R.Message = "customer cloud resources"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) UpdateCloudResource(ctx *gin.Context) {
	R := ResponseFormat{}

	cloudResourceUID := ctx.Param("id")
	if len(cloudResourceUID) == 0 {
		R.Error = append(R.Error, "id is required")
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	var i validators.UpdateCloudResourceInput
	if err := ctx.ShouldBindJSON(&i); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	// TODO: check to see if the name is available

	cloudResource, err := h.svc.GetCloudResourceByUID(cloudResourceUID)
	if err != nil {
		log.Println(err)
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	cloudResource.Name = i.Name
	cloudResource.Type = i.Type
	cloudResource.Region = i.Region

	if err := h.svc.UpdateCloudResource(cloudResource, cloudResource.ID); err != nil {
		log.Println(err)
		R.Message = messages.SomethingWentWrong
		ctx.JSON(h.Response(http.StatusInternalServerError, R))
		return
	}

	cloudResource, err = h.svc.GetCloudResourceByUID(cloudResourceUID)
	if err != nil {
		log.Println(err)
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	R.Data = cloudResource
	R.Message = "update cloud resources"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) DeleteCloudResource(ctx *gin.Context) {
	R := ResponseFormat{}

	R.Message = "delete cloud resources"
	ctx.JSON(h.Response(http.StatusOK, R))
}
