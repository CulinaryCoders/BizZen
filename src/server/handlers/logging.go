package handlers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"server/utils"
)

// RequestLoggingMiddleware logs the request type and URL of the inbound request and, if the request type is POST or PUT, also logs the JSON request type for debugging purposes
func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("INFO:  Received %s request for %s", request.Method, request.URL)

		//  Log JSON body of request for POST and PUT  requests for debugging purposes
		if request.Method == "POST" || request.Method == "PUT" {
			requestBodyBytes, err := io.ReadAll(request.Body)

			if err != nil {
				utils.RespondWithError(
					writer,
					http.StatusInternalServerError,
					err.Error())

				log.Panicf("ERROR:  %s", err.Error())
			}

			log.Printf("INFO:  Request body -- %s", string(requestBodyBytes))

			// Close body and reassign original body (only intended to be read once so needs to be reassigned)
			request.Body.Close()
			request.Body = io.NopCloser(bytes.NewBuffer(requestBodyBytes))
		}

		next.ServeHTTP(writer, request)
	})
}
