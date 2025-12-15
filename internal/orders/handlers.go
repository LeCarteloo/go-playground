package orders

import (
	"net/http"

	"go_playground/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service,
	}
}

func (handler *handler) ListOrders(writer http.ResponseWriter, request *http.Request) {
	orders, err := handler.service.ListOrders(request.Context())
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(writer, http.StatusOK, orders)
}
