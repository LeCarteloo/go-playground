package orders

import (
	"errors"
	"log"
	"net/http"

	"go_playground/internal/apperrors"
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

func (handler *handler) CreateOrder(writer http.ResponseWriter, request *http.Request) {
	var tempOrder createOrderParams

	if err := json.Read(request, &tempOrder); err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	createdOrder, err := handler.service.CreateOrder(request.Context(), tempOrder)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrProductNotFound):
			http.Error(writer, err.Error(), http.StatusNotFound)
		case errors.Is(err, apperrors.ErrInsufficientProductQuantity):
			http.Error(writer, err.Error(), http.StatusBadRequest)
		case errors.Is(err, apperrors.ErrInvalidCustomerID):
			http.Error(writer, err.Error(), http.StatusBadRequest)
		case errors.Is(err, apperrors.ErrNoOrderItems):
			http.Error(writer, err.Error(), http.StatusBadRequest)
		default:
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	json.Write(writer, http.StatusNotImplemented, createdOrder)
}
