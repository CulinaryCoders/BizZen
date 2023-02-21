package handlers

import (
	"net/http"
	"server/utils"
)

// TODO: Add comment documentation (func Authorize)
func (cookieHandler *CookieHandler) Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		session, _ := cookieHandler.Store.Get(request, "session")
		_, ok := session.Values["sessionID"]
		if !ok {
			utils.RespondWithError(writer, http.StatusUnauthorized, "Unauthorized.")
			return
		}

		handler.ServeHTTP(writer, request)
	}
}
