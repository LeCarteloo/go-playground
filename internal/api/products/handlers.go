package products

import (
	"errors"
	"net/http"
	"strconv"

	"go_playground/internal/apperrors"
	"go_playground/internal/json"
	"go_playground/internal/middleware"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (handler *handler) ListProducts(writer http.ResponseWriter, request *http.Request) {
	log := middleware.GetLogger(request.Context())

	products, err := handler.service.ListProducts(request.Context())
	if err != nil {
		log.Error("error listing products", "error", err)
		json.WriteError(writer, http.StatusInternalServerError, err)
		return
	}

	json.Write(writer, http.StatusOK, products)
}

func (handler *handler) GetProductById(writer http.ResponseWriter, request *http.Request) {
	log := middleware.GetLogger(request.Context())

	idParam := chi.URLParam(request, "productId")

	parsedIdParam, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Info("invalid product ID", "error", err)
		json.WriteError(writer, http.StatusBadRequest, apperrors.ErrInvalidProductID)
		return
	}

	product, err := handler.service.GetProductById(request.Context(), parsedIdParam)
	if err != nil {
		if errors.Is(err, apperrors.ErrProductNotFound) {
			json.WriteError(writer, http.StatusNotFound, err)
		} else {
			log.Error("error fetching product by ID", "error", err)
			json.WriteError(writer, http.StatusInternalServerError, err)
		}
		return
	}

	json.Write(writer, http.StatusOK, product)
}
