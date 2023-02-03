package middlewares

import (
	"encoding/json"
	"net/http"
	"server/controllers"
)

// TODO: Add comment documentation (func Authorize)
func Authorize(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		if request.Header["Token"] == nil {
			respondError(writer, http.StatusUnauthorized, "Token Not Found.")
			return
		}

		//var mySigningKey = []byte("secretkey")
		tokenString := request.Header["Token"][0]
		role, err := controllers.ValidateToken(tokenString)
		if err != nil {
			respondError(writer, http.StatusInternalServerError, err.Error())
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

// ? respondJSON duplicated across packages (present in both 'handlers' and 'middlewares')
// TODO: Consolidate respondJSON to a single package
// respondJSON makes the response with payload as json format
func respondJSON(writer http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	writer.Write([]byte(response))
}

// ? respondError duplicated across packages (present in both 'handlers' and 'middlewares')
// TODO: Consolidate respondJSON to a single package
// respondError makes the error response with payload as json format
func respondError(writer http.ResponseWriter, code int, message string) {
	respondJSON(
		writer,
		code,
		map[string]string{"error": message},
	)
}
