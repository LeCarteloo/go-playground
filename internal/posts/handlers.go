package posts

import (
	"go_playground/internal/json"
	"log"
	"net/http"
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
	err := handler.service.ListPosts(request.Context())

	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	posts := []string{"Post 1", "Post 2", "Post 3"}

	json.Write(writer, http.StatusOK, posts)
}
