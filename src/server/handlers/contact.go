package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/utils"
	_ "time"
)

/*
*Description*

func CreateAddress

Creates a new address record in the database.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  POST

	Route:	/address

	Body:
		Format: JSON

		Required fields:

			address1  <string>

				First line of the street address

			city  <string>

				City

			state  <string>

				Two letter capitalized state abbreviation

			zip  <string>

				Five digit zip code or nine digit postal code

		Optional fields:

			address2  <string>

				Second line of the street address

*Example request(s)*

	POST /address
	{
		"address1":"123 Test Address Rd",
		"address2":"",
		"city":"Gainesville",
		"state":"FL",
		"zip":"12345"
	}

*Response format*

	Success:

		HTTP/1.1 201 Created
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": null,
			"address1":"123 Test Address Rd",
			"address2":"",
			"city":"Gainesville",
			"state":"FL",
			"zip": "12345"
		}

	Failure:

		-- Case = Bad request body
		HTTP/1.1 400 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Case = Database operation error
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) CreateAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&address); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	returnedRecords, err := address.Create(app.AppDB)
	createdAddress := returnedRecords["address"]
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusCreated,
		createdAddress)
}

/*
*Description*

func GetAddress

Get address record from the database by ID.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/address/{id}

	Body:

		None

*Example request(s)*

	GET /address/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": null,
			"address1":"123 Test Address Rd",
			"address2":"",
			"city":"Gainesville",
			"state":"FL",
			"zip": "12345"
		}

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		HTTP/1.1 404 Resource Not Found
		Content-Type: application/json

		{
			"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) GetAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}
	addressID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnRecords, err := address.Get(app.AppDB, addressID)
	returnedAddress := returnRecords["address"]
	if err != nil {
		var errorMessage string = fmt.Sprintf("Address ID (%d) does not exist in the database.\n%s", addressID, err)

		utils.RespondWithError(
			writer,
			http.StatusNotFound,
			errorMessage)

		log.Panicf(errorMessage)

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		returnedAddress)
}

/*
*Description*

func UpdateAddress

Updates the address record associated with the specified address ID in the database.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "address2": "").

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:   PUT

	Route:  /address/{id}

	Body:
		Format: JSON

		Required fields:

			N/A  --  At least one field should be present in the request body, but no fields are specifically required to be present in the request body.

		Optional fields:

			address1  <string>

				First line of the street address

			address2  <string>

				Second line of the street address

			city  <string>

				City

			state  <string>

				Two letter capitalized state abbreviation

			zip  <string>

				Five digit zip code or nine digit postal code

*Example request(s)*

	PUT /address/123456
	{
		"address1":"789 Updated Address Blvd",
		"city":"Orlando",
		"zip":"45678"
	}

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2022-07-11T01:23:45.6789012-14:25",
			"DeletedAt": null,
			"address1":"789 Updated Address Blvd",
			"address2":"",
			"city":"Orlando",
			"state":"FL",
			"zip": "45678"
		}

	Failure:
		-- Case = Bad request body or missing/misformatted ID in request URL
		HTTP/1.1 400 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Case = Database operation error
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) UpdateAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}
	addressID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	var updates map[string]interface{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updates); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	returnedRecords, err := address.Update(app.AppDB, addressID, updates)
	updatedAddress := returnedRecords["address"]
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		updatedAddress)
}

/*
*Description*

func DeleteAddress

Delete an address record from the database by address ID if the ID exists in the database.

Deleted address record is returned in the response body if the operation is sucessful.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	DELETE

	Route:	/address/{id}

	Body:

		None

*Example request(s)*

	DELETE /address/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": "2022-07-11T01:23:45.6789012-14:25",
			"address1":"123 Test Address Rd",
			"address2":"",
			"city":"Gainesville",
			"state":"FL",
			"zip": "12345"
		}

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Case = Database operation error
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) DeleteAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}
	addressID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedRecords, err := address.Delete(app.AppDB, addressID)
	deletedAddress := returnedRecords["address"]
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		deletedAddress)

}
