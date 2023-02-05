package handlers

import "net/http"

// TODO: Add comment documentation (func Customer)
func (db Handler) Customer(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("Role") != "customer" {
		RespondError(
			writer,
			http.StatusBadRequest,
			"Invalid role.",
		)

		return
	}

	writer.Write([]byte("Welcome, Customer."))
}
