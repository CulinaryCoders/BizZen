package utils

import (
	"encoding/json"
	"net/http"
)

/*
RespondWithJSON marshals the payload into JSON format and returns the HTTP response

Parameters:

- writer (http.ResponseWriter): Response writer used to write the response to the HTTP client.
- code (int): Status code used to set the HTTP status code for the response.
- payload (interface{}): Payload used to serialize to JSON format and send to the HTTP client as the response body.

Returns:

- This function does not return any values.

Description:

The function uses the json.Marshal function to encode the payload to a JSON byte slice. If there is an error during the encoding process, the function will return an HTTP 500 Internal Server Error response with an error message.

The function sets the response header's content type to application/json and writes the serialized JSON payload to the response writer.
*/
func RespondWithJSON(writer http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write([]byte(response))
}

/*
RespondWithError takes in an error code and error message and returns the HTTP response

Parameters:

* writer (http.ResponseWriter): Response writer that will be used to send the error response.
* code (int): HTTP status code to be used in the error response.
* message (string): Error message to be sent in the response.
Returns:
* The RespondWithError function does not return any values.

Description:
RespondWithError function takes a http.ResponseWriter, a http status code and a message as input parameters. It formats the error message as a JSON object with a "error" field containing the message, and writes it to the ResponseWriter with the given status code.
*/
func RespondWithError(writer http.ResponseWriter, code int, message string) {
	RespondWithJSON(
		writer,
		code,
		map[string]string{"error": message},
	)
}
