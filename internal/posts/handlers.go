package posts

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service,
	}
}

func (handler *handler) ListPosts(writer http.ResponseWriter, request *http.Request) {

	posts := []string{"Hello", "World"}

	json.NewEncoder(writer).Encode(posts)
}
