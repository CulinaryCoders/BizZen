package handlers

import "net/http"

// TODO: Add comment documentation (func Customer)
func (db Handler) Business(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("Role") != "business" {
		RespondError(
			writer,
			http.StatusBadRequest,
			"Invalid role.",
		)

		return
	}

	RespondJSON(
		writer,
		http.StatusOK,
		"Welcome, Business Owner.",
	)
}
