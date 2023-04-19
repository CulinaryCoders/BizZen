package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
*Description*

func RespondWithJSON

Marshals the payload into JSON format and returns the HTTP response.

The method uses the json.Marshal method to encode the payload to a JSON byte slice. If there is an error during the encoding process, the method will return an HTTP 500 Internal Server Error response with an error message.

The method sets the response header's content type to application/json and writes the serialized JSON payload to the response writer.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

	payload <interface{}>

		Payload used to serialize to JSON format and send to the HTTP client as the response body.

*Returns*

	None
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
*Description*

func RespondWithJSON

Takes a http.ResponseWriter, a HTTP status code and a message as input parameters. It formats the error message as a JSON object with a "error" field containing the message, and writes it to the ResponseWriter with the given status code.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	code  <int>

		The HTTP status code to be used in the error response.

	message  <string>

		Error message to be sent in the response.

*Returns*

	None
*/
func RespondWithError(writer http.ResponseWriter, code int, message string) {
	RespondWithJSON(
		writer,
		code,
		map[string]string{"error": message},
	)
}

/*
*Description*

func ParseRequestID

Helper method to parse the "id" variable present in the request and convert it to an unsigned integer.

*Parameters*

	request  <*http.Request>

		The HTTP request.

*Returns*

	_  <uint>

		The "id" variable value that has been converted to uint format.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func ParseRequestID(request *http.Request) (uint, error) {
	userId := mux.Vars(request)["id"]
	convertedToUint64, err := strconv.ParseUint(userId, 10, 64)
	return uint(convertedToUint64), err
}

/*
*Description*

func ParseRequestIDField

Helper method to parse the specified ID variable present in the request and convert it to an unsigned integer.

Offers a more generic alternative to the ParseRequestID method, in instances where there are multiple ID variables in the route
or when a route has an ID variable with a key other than "id".

*Parameters*

	request  <*http.Request>

		The HTTP request.

	idFieldKey  <string>

		The ID variable key used for the request's route.

*Returns*

	_  <uint>

		The specified ID variable's value that has been converted to uint format.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func ParseRequestIDField(request *http.Request, idFieldKey string) (uint, error) {
	id := mux.Vars(request)[idFieldKey]
	convertedToUint64, err := strconv.ParseUint(id, 10, 64)
	return uint(convertedToUint64), err
}

/*
*Description*

func DecodeJSON

Helper method to unmarshal the request body into the provided gorm.Model object.

*Parameters*

	request  <*http.Request>

		The HTTP request.

	modelObj  <interface{}>

		The gorm.Model object that the request body JSON will be unmarshalled into.

*Returns*

	_  <uint>

		The updated gorm.Model object that the request's JSON body was unmarshalled into.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func DecodeJSON(request *http.Request, modelObj interface{}) (interface{}, error) {
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(modelObj); err != nil {
		return modelObj, err
	}
	defer request.Body.Close()
	return modelObj, nil
}
