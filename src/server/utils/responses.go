package utils

import (
	"encoding/json"
	"net/http"
)

// RespondJSON makes the response with payload as json format
func RespondJSON(writer http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func RespondError(writer http.ResponseWriter, code int, message string) {
	RespondJSON(
		writer,
		code,
		map[string]string{"error": message},
	)
}
