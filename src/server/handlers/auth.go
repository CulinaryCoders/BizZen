package handlers

import (
	"net/http"
	"server/utils"
)

// Authorize is a middleware function that uses Gorilla sessions to verify that the user is authenticated before calling the next HTTP handler function in the chain.
//
// Parameters:
//   - next http.HandlerFunc: the next HTTP handler function in the chain to call if authentication is successful.
//
// Returns:
//   - http.HandlerFunc: a new HTTP handler function that performs authentication and calls the next handler function if authentication is successful.
//
// Behavior:
//
//	The middleware function first retrieves the session cookie from the HTTP request using the Gorilla sessions package. If the session cookie is not found, or if the session does not contain a "sessionID" value, the function responds with a 401 Unauthorized error and stops the chain of HTTP handlers.
//
//	If the user is authenticated, the function calls the next HTTP handler function in the chain with the same http.ResponseWriter and *http.Request objects passed to it.
func (h *Handler) Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		session, _ := h.Store.Get(request, "session")
		_, ok := session.Values["sessionID"]
		if !ok {
			utils.RespondWithError(writer, http.StatusUnauthorized, "Unauthorized.")
			return
		}

		next.ServeHTTP(writer, request)
	}
}
