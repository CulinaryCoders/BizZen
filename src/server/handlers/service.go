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

			office_id  <uint>

				ID of Office record Service is associated with

			name  <string>

				Name of the service

			desc  <string>

				Description of the service

		Optional fields:

			N/A (None)

*Example request(s)*

	POST /service
	{
		"office_id":123
		"name":"Yoga class",
		"desc":"30 minute beginner yoga class"
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
			"office_id":123
			"name":"Yoga class",
			"desc":"30 minute beginner yoga class"
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
			"office_id":123,
			"name":"Yoga class",
			"desc":"30 minute beginner yoga class"
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
		var errorMessage string = fmt.Sprintf("Service ID (%d) does not exist in the database.\n%s", serviceID, err)

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

			office_id  <uint>

				ID of Office record Service is associated with

			name  <string>

				Name of the service

			desc  <string>

				Description of the service

*Example request(s)*

	PUT /service/123456
	{
		"office_id":456,
		"name":"Personal training session",
		"desc":"1 hour personal training session with qualified trainer"
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
			"office_id":456,
			"name":"Personal training session",
			"desc":"1 hour personal training session with qualified trainer"
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
			"office_id":123,
			"name":"Yoga class",
			"desc":"30 minute beginner yoga class"
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

// TODO:  Update response info in docstring with correct fields/values (func CreateServiceOffering)
/*
*Description*

func CreateServiceOffering

Creates a new ServiceOffering record in the database.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  POST

	Route:	/service-offering

	Body:
		Format: JSON

		Required fields:

			service_id  <uint>

				ID of the Service record that the ServiceOffering is associated with

			start_date  <time.Time>

				ServiceOffering start date

			end_date  <time.Time>

				ServiceOffering end date

			booking_length  <uint>

				Length of time appointment booking is for (in minutes)

		Optional fields:

			staff_id  <uint>

				ID of Staff member that ServiceOffering is associated with

			resource_id  <uint>

				ID of Resource that ServiceOffering is associated with

			price  <uint>

				Price (in cents) for the service being offered

			cancel_fee  <uint>

				Fee (in cents) for cancelling appointment after minimum notice cutoff

			max_consec_bookings  <uint>

				Max number of consecutive appointments customers can book

			min_cancel_notice  <uint>

				Minimum number of hours appointment cancellation must be made in order to avoid cancellation fee. (null if not applicable, 0 if cancellation fee is always applied)

			min_time_betw_clients  <uint>

				Length of time between appointments for differing clients

*Example request(s)*

	POST /service-offering
	{
		"service_id":123,
		"start_date":"2021-01-01T00:00:00.0000000-00:00",
		"end_date":"2023-01-01T00:00:00.0000000-00:00",
		"booking_length":30,
		"price":10000
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
			"service_id":123,
			"start_date":"2021-01-01T00:00:00.0000000-00:00",
			"end_date":"2023-01-01T00:00:00.0000000-00:00",
			"booking_length":30,
			"price":10000,
			"staff_id":,
			"resource_id":,
			"cancel_fee":,
			"max_consec_bookings":,
			"min_cancel_notice":,
			"min_time_betw_clients":
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
func (app *Application) CreateServiceOffering(writer http.ResponseWriter, request *http.Request) {
	serviceOffering := models.ServiceOffering{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&serviceOffering); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	returnedRecords, err := serviceOffering.Create(app.AppDB)
	createdServiceOffering := returnedRecords["service_offering"]
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
		createdServiceOffering)
}

// TODO:  Update response info in docstring with correct fields/values (func GetServiceOffering)
/*
*Description*

func GetServiceOffering

Get ServiceOffering record from the database by ID.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/service-offering/{id}

	Body:

		None

*Example request(s)*

	GET /service-offering/123456

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
func (app *Application) GetServiceOffering(writer http.ResponseWriter, request *http.Request) {
	serviceOffering := models.ServiceOffering{}
	serviceOfferingID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedServiceOffering, err := serviceOffering.Get(app.AppDB, serviceOfferingID)
	if err != nil {
		var errorMessage string = fmt.Sprintf("ServiceOffering ID (%d) does not exist in the database.\n%s", serviceOfferingID, err)

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
		returnedServiceOffering)
}

// TODO:  Update response info in docstring with correct fields/values (func UpdateServiceOffering)
/*
*Description*

func UpdateServiceOffering

Updates the ServiceOffering record associated with the specified ServiceOffering ID in the database.

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

	Route:  /service-offering/{id}

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

	PUT /service-offering/123456
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
func (app *Application) UpdateServiceOffering(writer http.ResponseWriter, request *http.Request) {
	serviceOffering := models.ServiceOffering{}
	serviceOfferingID, err := utils.ParseRequestID(request)

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

	returnedRecords, err := serviceOffering.Update(app.AppDB, serviceOfferingID, updates)
	updatedServiceOffering := returnedRecords["service_offering"]
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
		updatedServiceOffering)
}

// TODO:  Update response info in docstring with correct fields/values (func DeleteServiceOffering)
/*
*Description*

func DeleteServiceOffering

Delete an ServiceOffering record from the database by ServiceOffering ID if the ID exists in the database.

Deleted ServiceOffering record is returned in the response body if the operation is sucessful.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	DELETE

	Route:	/service-offering/{id}

	Body:

		None

*Example request(s)*

	DELETE /service-offering/123456

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
func (app *Application) DeleteServiceOffering(writer http.ResponseWriter, request *http.Request) {
	serviceOffering := models.ServiceOffering{}
	serviceOfferingID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedRecords, err := serviceOffering.Delete(app.AppDB, serviceOfferingID)
	deletedServiceOffering := returnedRecords["service_offering"]
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
		deletedServiceOffering)

}
