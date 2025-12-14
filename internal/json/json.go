package json

import (
	"encoding/json"
	"net/http"
)

func Write(writer http.ResponseWriter, status int, data any) {
	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}
