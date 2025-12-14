package posts

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

func (handler *handler) ListPosts(writer http.ResponseWriter, request *http.Request) {
	posts, err := handler.service.ListPosts(request.Context())
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(writer, http.StatusOK, posts)
}

func (handler *handler) GetPostById(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "postId")

	parsedIdParam, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(writer, "invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := handler.service.GetPostById(request.Context(), parsedIdParam)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Fetched post:", post)

	json.Write(writer, http.StatusOK, post)
}
