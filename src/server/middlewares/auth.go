package middlewares

import (
	"net/http"
	"server/controllers"
	"server/handlers"
)

// TODO: Add comment documentation (func Authorize)
func Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Header["Token"] == nil {
			handlers.RespondError(writer, http.StatusUnauthorized, "Token Not Found.")
			return
		}

		//var mySigningKey = []byte("secretkey")
		tokenString := request.Header["Token"][0]
		role, err := controllers.ValidateToken(tokenString)
		if err != nil {
			handlers.RespondError(writer, http.StatusInternalServerError, err.Error())
			return
		}

		if role == "customer" {
			request.Header.Set("Role", "customer")
			handler.ServeHTTP(writer, request)
			return

		} else if role == "user" {
			request.Header.Set("Role", "user")
			handler.ServeHTTP(writer, request)
			return

		}

	}
}
