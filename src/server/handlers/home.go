package handlers

import (
	"encoding/json"
	"net/http"
)

func (db Handler) Home(writer http.ResponseWriter, request *http.Request) {
	// your code here
	writer.Write([]byte("Welcome to the Matrix!"))
	json.NewEncoder(writer).Encode("Product Not Found!")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}
