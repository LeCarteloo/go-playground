package orders

import (
	"errors"
	"log/slog"
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
		slog.Error("error listing orders", "error", err)
		json.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	json.Write(writer, http.StatusOK, orders)
}

func (handler *handler) CreateOrder(writer http.ResponseWriter, request *http.Request) {
	var tempOrder createOrderParams

	if err := json.Read(request, &tempOrder); err != nil {
		slog.Error("error with decoding body", "error", err)
		json.WriteError(writer, http.StatusBadRequest, err)
		return
	}

	createdOrder, err := handler.service.CreateOrder(request.Context(), tempOrder)
	if err != nil {

		slog.Error("error creating order", "error", err)
		switch {
		case errors.Is(err, apperrors.ErrProductNotFound):
			json.WriteError(writer, http.StatusNotFound, err)
		case errors.Is(err, apperrors.ErrInsufficientProductQuantity):
			json.WriteError(writer, http.StatusBadRequest, err)
		case errors.Is(err, apperrors.ErrInvalidCustomerID):
			json.WriteError(writer, http.StatusBadRequest, err)
		case errors.Is(err, apperrors.ErrNoOrderItems):
			json.WriteError(writer, http.StatusBadRequest, err)
		default:
			json.WriteError(writer, http.StatusInternalServerError, err)
		}
		return
	}

	json.Write(writer, http.StatusNotImplemented, createdOrder)
}
