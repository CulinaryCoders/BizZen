package handlers

import "net/http"

func (h handler) Customer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "customer" {
		respondError(w, http.StatusBadRequest, "Invalid role.")
		return
	}
	w.Write([]byte("Welcome, Customer."))
}
