package handlers

import (
	"net/http"
	"server/config"
	"server/utils"
)

// TODO: Add comment documentation (func Authorize)
func Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Header["Token"] == nil {
			utils.RespondWithError(writer, http.StatusUnauthorized, "Token Not Found.")
			return
		}

		//var mySigningKey = []byte("secretkey")
		tokenString := request.Header["Token"][0]
		role, err := ValidateToken(tokenString, config.AppConfig.GetSigningKey())
		if err != nil {
			utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
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
