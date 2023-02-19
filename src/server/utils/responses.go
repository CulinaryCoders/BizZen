package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON marshals the payload into JSON format and returns the HTTP response
func RespondWithJSON(writer http.ResponseWriter, status int, payload interface{}) {
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

// RespondWithError takes in an error code and error message and returns the HTTP response
func RespondWithError(writer http.ResponseWriter, code int, message string) {
	RespondWithJSON(
		writer,
		code,
		map[string]string{"error": message},
	)
}
