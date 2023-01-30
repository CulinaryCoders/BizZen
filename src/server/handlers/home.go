package handlers

import (
	"encoding/json"
	"net/http"
)

func (h handler) Home(w http.ResponseWriter, r *http.Request) {
	// your code here
	w.Write([]byte("Welcome to the Matrix!"))
	json.NewEncoder(w).Encode("Product Not Found!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
