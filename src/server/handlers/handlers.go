package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) Handler {
	return Handler{db}
}

// respondJSON makes the response with payload as json format
func respondJSON(writer http.ResponseWriter, status int, payload interface{}) {
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

// respondError makes the error response with payload as json format
func respondError(writer http.ResponseWriter, code int, message string) {
	respondJSON(
		writer,
		code,
		map[string]string{"error": message},
	)
}
