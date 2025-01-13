package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/babyfaceeasy/crims/internal/messages"
	"github.com/babyfaceeasy/crims/internal/validators"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		R.Message = messages.NotFound
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	err = h.svc.AddCloudResourcesToCustomer(customer.ID, i.Resources)
	if err != nil {
		log.Println(err)
		R.Message = messages.SomethingWentWrong
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
	if err != nil ||  customer == nil{
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

	// get cloud resource uid
	id := ctx.Param("id")
	if id == "" {
		R.Message = "cloud resource id is required"
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

	cloudResource, err := h.svc.GetCloudResourceByUID(id)
	if err != nil {
		log.Println(err)
		R.Message = messages.SomethingWentWrong
		ctx.JSON(h.Response(http.StatusInternalServerError, R))
		return
	}

	// check if a change is been made to the name and confirm if the name is available
	if cloudResource.Name != i.Name {
		// check to see if the new name is available
		available, err := h.svc.IsCloudResourceNameAvailable(i.Name)
		if err != nil {
			log.Println(err)
			R.Message = messages.SomethingWentWrong
			ctx.JSON(h.Response(http.StatusInternalServerError, R))
			return
		}

		if !available {
			R.Message = messages.CloudResourceNameAlreadyInUse
			ctx.JSON(h.Response(http.StatusBadRequest, R))
			return
		}
	}

	// update the cloud resource
	cloudResource.Name = i.Name
	cloudResource.Type = i.Type
	cloudResource.Region = i.Region
	if err := h.svc.UpdateCloudResource(cloudResource, cloudResource.ID); err != nil {
		log.Println(err)
		R.Message = messages.SomethingWentWrong
		ctx.JSON(h.Response(http.StatusInternalServerError, R))
		return
	}

	R.Data = cloudResource
	R.Message = "cloud resource update was successful"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) DeleteCloudResource(ctx *gin.Context) {
	R := ResponseFormat{}

	id := ctx.Param("id")
	if id == "" {
		R.Message = "cloud resource id is required"
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	cloudResource, err := h.svc.GetCloudResourceByUID(id)
	if err != nil {
		// TODO: create your own errors here to make it more solid
		if errors.Is(err, gorm.ErrRecordNotFound) {
			R.Message = messages.CloudResourceNotFound
			ctx.JSON(h.Response(http.StatusNotFound, R))
			return
		}
		log.Println(err)
		R.Message = messages.SomethingWentWrong
		ctx.JSON(h.Response(http.StatusInternalServerError, R))
		return
	}

	if err := h.svc.DeleteCloudResource(cloudResource.ID); err != nil {
		log.Println(err)
		R.Message = messages.SomethingWentWrong
		ctx.JSON(h.Response(http.StatusInternalServerError, R))
		return
	}

	R.Message = "cloud resource deleted successfully"
	ctx.JSON(h.Response(http.StatusOK, R))
}
