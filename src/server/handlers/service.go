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

func CreateService

Creates a new service record in the database.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  POST

	Route:	/service

	Body:
		Format: JSON

		Required fields:

			business_id  <uint>

				ID of Business record Service is associated with

			name  <string>

				Name of the service

			desc  <string>

				Description of the service

			start_date_time  <time.Time>

				Date/time that service is scheduled to start

			length <uint>

				Length of time in minutes that the service will take

			capacity <uint>

				Number of users that can sign up for the service

			price <uint>

				Price (in cents) for the service being offered

		Optional fields:

			cancel_fee <uint>

				Fee (in cents) for cancelling appointment after minimum notice cutoff

*Example request(s)*

	POST /service
	{
		"business_id":123
		"name":"Yoga class",
		"desc":"30 minute beginner yoga class",
		"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
		"length":30,
		"capacity":20,
		"price":2000
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
			"business_id":123
			"name":"Yoga class",
			"desc":"30 minute beginner yoga class",
			"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
			"length":30,
			"capacity":20,
			"price":2000,
			"cancel_fee":0
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
func (app *Application) CreateService(writer http.ResponseWriter, request *http.Request) {
	service := models.Service{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&service); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	returnedRecords, err := service.Create(app.AppDB)
	createdService := returnedRecords["service"]
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
		createdService)
}

/*
*Description*

func GetService

Get service record from the database by ID.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/service/{id}

	Body:

		None

*Example request(s)*

	GET /service/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": null,
			"business_id":123
			"name":"Yoga class",
			"desc":"30 minute beginner yoga class",
			"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
			"length":30,
			"capacity":20,
			"price":2000,
			"cancel_fee":0
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
func (app *Application) GetService(writer http.ResponseWriter, request *http.Request) {
	service := models.Service{}
	serviceID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedService, err := service.Get(app.AppDB, serviceID)
	if err != nil {
		var errorMessage string = fmt.Sprintf("Service ID (%d) does not exist in the database.  [%s]", serviceID, err)

		utils.RespondWithError(
			writer,
			http.StatusNotFound,
			errorMessage)

		log.Printf("ERROR:  %s", errorMessage)

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		returnedService)
}

/*
*Description*

func UpdateService

Updates the service record associated with the specified service ID in the database.

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

	Route:  /service/{id}

	Body:
		Format: JSON

		Required fields:

			N/A  --  At least one field should be present in the request body, but no fields are specifically required to be present in the request body.

		Optional fields:

			business_id  <uint>

				ID of Business record Service is associated with

			name  <string>

				Name of the service

			desc  <string>

				Description of the service

			start_date_time  <time.Time>

				Date/time that service is scheduled to start

			length <uint>

				Length of time in minutes that the service will take

			capacity <uint>

				Number of users that can sign up for the service

			price <uint>

				Price (in cents) for the service being offered

			cancel_fee <uint>

				Fee (in cents) for cancelling appointment after minimum notice cutoff

*Example request(s)*

	PUT /service/123456
	{
		"price":2500,
		"cancel_fee":1000,
	}

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-02-13T04:20:12.6789012-05:00",
			"DeletedAt": null,
			"business_id":123
			"name":"Yoga class",
			"desc":"30 minute beginner yoga class",
			"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
			"length":30,
			"capacity":20,
			"price":2500,
			"cancel_fee":1000
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
func (app *Application) UpdateService(writer http.ResponseWriter, request *http.Request) {
	service := models.Service{}
	serviceID, err := utils.ParseRequestID(request)

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

	returnedRecords, err := service.Update(app.AppDB, serviceID, updates)
	updatedService := returnedRecords["service"]
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
		updatedService)
}

/*
*Description*

func DeleteService

Delete an service record from the database by service ID if the ID exists in the database.

Deleted service record is returned in the response body if the operation is sucessful.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	DELETE

	Route:	/service/{id}

	Body:

		None

*Example request(s)*

	DELETE /service/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": "2022-07-11T01:23:45.6789012-14:25",
			"business_id":123
			"name":"Yoga class",
			"desc":"30 minute beginner yoga class",
			"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
			"length":30,
			"capacity":20,
			"price":2000,
			"cancel_fee":0
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
func (app *Application) DeleteService(writer http.ResponseWriter, request *http.Request) {
	service := models.Service{}
	serviceID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedRecords, err := service.Delete(app.AppDB, serviceID)
	deletedService := returnedRecords["service"]
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
		deletedService)

}

// TODO:  Add documentation (func GetListOfEnrolledUsers)
func (app *Application) GetListOfEnrolledUsers(writer http.ResponseWriter, request *http.Request) {
	// log.Print("Undefined route handler requested  --  GetListOfEnrolledUsers")
	service := models.Service{}
	serviceID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	var users []models.User
	users, err = service.GetUsers(app.AppDB, serviceID)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusNotFound,
			err.Error())

		log.Printf("ERROR:  %s", err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		users)
}

// TODO:  Add documentation (func GetEnrolledUsersCount)
func (app *Application) GetEnrolledUsersCount(writer http.ResponseWriter, request *http.Request) {
	// log.Print("Undefined route handler requested  --  GetEnrolledUsersCount")
	service := models.Service{}
	serviceID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	var users []models.User
	users, err = service.GetUsers(app.AppDB, serviceID)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusNotFound,
			err.Error())

		log.Printf("ERROR:  %s", err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		len(users))
}
