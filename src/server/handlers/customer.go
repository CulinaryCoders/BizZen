package handlers

import "net/http"

func (db Handler) Customer(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("Role") != "customer" {
		respondError(writer, http.StatusBadRequest, "Invalid role.")
		return
	}
	writer.Write([]byte("Welcome, Customer."))
}
