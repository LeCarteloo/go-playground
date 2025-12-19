package orders

import (
	"errors"
	"net/http"

	"go_playground/internal/apperrors"
	"go_playground/internal/json"
	"go_playground/internal/middleware"
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
	log := middleware.GetLogger(request.Context())

	orders, err := handler.service.ListOrders(request.Context())
	if err != nil {
		log.Error("error listing orders", "error", err)
		json.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	json.Write(writer, http.StatusOK, orders)
}

func (handler *handler) CreateOrder(writer http.ResponseWriter, request *http.Request) {
	log := middleware.GetLogger(request.Context())

	var tempOrder createOrderParams

	if err := json.Read(request, &tempOrder); err != nil {
		log.Warn("error with decoding body", "error", err)
		json.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	createdOrder, err := handler.service.CreateOrder(request.Context(), tempOrder)
	if err != nil {

		switch {
		case errors.Is(err, apperrors.ErrProductNotFound):
			json.WriteError(writer, http.StatusNotFound, err)
		case errors.Is(err, apperrors.ErrInsufficientProductQuantity),
			errors.Is(err, apperrors.ErrInvalidCustomerID),
			errors.Is(err, apperrors.ErrNoOrderItems):
			json.WriteError(writer, http.StatusBadRequest, err)
		default:
			log.Error("failed to create order", "error", err)
			json.WriteError(writer, http.StatusInternalServerError, err)
		}
		return
	}

	json.Write(writer, http.StatusCreated, createdOrder)
}
