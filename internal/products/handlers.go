package products

import (
	"log"
	"net/http"
	"strconv"

	"go_playground/internal/json"

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
	products, err := handler.service.ListProducts(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(writer, http.StatusOK, products)
}

func (handler *handler) GetProductById(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "productId")

	parsedIdParam, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, "invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := handler.service.GetProductById(request.Context(), parsedIdParam)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Fetched product:", product)

	json.Write(writer, http.StatusOK, product)
}
