package handlers

import (
	"log"
	"net/http"
)

func (h Handler) Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Header["Token"] == nil {
			RespondError(writer, http.StatusUnauthorized, "Token Not Found.")
			return
		}

		//var mySigningKey = []byte("secretkey")
		tokenString := request.Header["Token"][0]
		role, err := ValidateToken(tokenString)
		if err != nil {
			RespondError(writer, http.StatusInternalServerError, err.Error())
			return
		}

		log.Println(role)

		if role == "customer" {
			request.Header.Set("Role", "customer")
			handler.ServeHTTP(writer, request)
			return

		} else if role == "business" {
			request.Header.Set("Role", "business")
			handler.ServeHTTP(writer, request)
			return

		}
		RespondError(writer, http.StatusInternalServerError, "No role found.")

	}
}

func (h Handler) AuthMidleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := TokenValid(request)
		if err != nil {
			RespondError(writer, http.StatusUnauthorized, "Token Not Found.")
			return
		}

		handler.ServeHTTP(writer, request)
	}
}
