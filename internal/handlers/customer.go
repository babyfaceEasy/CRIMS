package handlers

import (
	"log"
	"net/http"

	"github.com/babyfaceeasy/crims/internal/messages"
	"github.com/babyfaceeasy/crims/internal/validators"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetCustomers(ctx *gin.Context) {
	R := ResponseFormat{}

	R.Message = "list of customers"
	ctx.JSON(h.Response(http.StatusOK, R))
}

func (h Handler) CreateCustomer(ctx *gin.Context) {
	R := ResponseFormat{}

	var i validators.CreateCustomerInput
	if err := ctx.ShouldBindJSON(&i); err != nil {
		R.Error = append(R.Error, err.Error())
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusBadRequest, R))
		return
	}

	emailTaken, err := h.svc.IsEmailTaken(i.Email)
	if err != nil {
		log.Println(err)
		R.Message = messages.SomethingWentWrong
		ctx.JSON(h.Response(http.StatusInternalServerError, R))
		return
	}

	if emailTaken {
		R.Error = append(R.Error, "email is taken")
		R.Message = messages.ValidationFailed
		ctx.JSON(h.Response(http.StatusUnprocessableEntity, R))
		return
	}

	err = h.svc.AddCustomer(i.Name, i.Email)
	if err != nil {
		log.Printf("error occurred in CreateCustomer: %s", err)
		R.Message = messages.SomethingWentWrong
		ctx.JSON(h.Response(http.StatusInternalServerError, R))
		return
	}

	R.Message = messages.CustomerCreated
	ctx.JSON(h.Response(http.StatusCreated, R))
}
