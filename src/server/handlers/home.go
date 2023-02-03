package handlers

import (
	"encoding/json"
	"net/http"
)

// TODO: Add comment documentation (func Home)
func (db Handler) Home(writer http.ResponseWriter, request *http.Request) {
	// your code here
	writer.Write([]byte("Welcome to BizZen!"))
	json.NewEncoder(writer).Encode("Resource Not Found!")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
