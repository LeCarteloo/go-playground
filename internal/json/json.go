package json

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Write(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(data)
}

func WriteError(writer http.ResponseWriter, status int, error error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(ErrorResponse{Message: error.Error()})
}

func Read[T any](request *http.Request, data *T) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}
