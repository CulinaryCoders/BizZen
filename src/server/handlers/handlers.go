package handlers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

// TODO: Add comment documentation (type Handler)
type Handler struct {
	DB *gorm.DB
}

// TODO: Add comment documentation (func NewHandler)
func NewHandler(db *gorm.DB) Handler {
	return Handler{db}
}

// ? respondJSON duplicated across packages (present in both 'handlers' and 'middlewares')
// TODO: Consolidate respondJSON to a single package
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

// ? respondError duplicated across packages (present in both 'handlers' and 'middlewares')
// TODO: Consolidate respondJSON to a single package
// respondError makes the error response with payload as json format
func respondError(writer http.ResponseWriter, code int, message string) {
	respondJSON(
		writer,
		code,
		map[string]string{"error": message},
	)
}
