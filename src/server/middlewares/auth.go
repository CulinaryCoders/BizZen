package middlewares

import (
	"encoding/json"
	"net/http"
	"server/controllers"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			respondError(w, http.StatusUnauthorized, "Token Not Found.")
			return
		}

		//var mySigningKey = []byte("secretkey")
		tokenString := r.Header["Token"][0]
		role, err := controllers.ValidateToken(tokenString)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if role == "customer" {
			r.Header.Set("Role", "customer")
			handler.ServeHTTP(w, r)
			return

		} else if role == "user" {
			r.Header.Set("Role", "user")
			handler.ServeHTTP(w, r)
			return

		}

	}
}

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
