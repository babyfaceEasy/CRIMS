package handlers

import (
	"net/http"

	"github.com/babyfaceeasy/crims/internal/messages"
	"github.com/babyfaceeasy/crims/internal/services"
)

type Handler struct {
	svc services.ServiceInterface
}

func NewHandler(service services.ServiceInterface) Handler {
	return Handler{svc: service}
}

// Handles the structure for returning Response
type ResponseFormat struct {
	Status  bool     `json:"status"`
	Data    any      `json:"data"`
	Error   []string `json:"error"`
	Message string   `json:"message"`
}

// Response create new instance of ResponseFormat
func (h Handler) Response(code int, res ResponseFormat) (int, ResponseFormat) {
	if code == 0 {
		code = 200
	}

	if !res.Status {
		res.Status = code < http.StatusBadRequest
	}

	if res.Data == nil {
		res.Data = make(map[string]interface{})
	}
	if res.Error == nil {
		res.Error = make([]string, 0)
	}
	if res.Message == "" {
		res.Message = messages.OperationWasSuccessful
		if code == http.StatusNotFound {
			res.Message = messages.NotFound
		} else if code >= http.StatusBadRequest {
			res.Message = messages.SomethingWentWrong
		}
	}

	return code, res
}
