package handlers

import "net/http"

func (h handler) Customer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "customer" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Customer."))
}
