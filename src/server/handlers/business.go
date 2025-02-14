package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/utils"
)

/*
*Description*

func CreateBusiness

Creates a new business record in the database and returns them as JSON in the response body.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  POST

	Route:	/business

	Body:
		Format: JSON

		Required fields:

			owner_id  <uint>

				ID of the User account that created/owns this business record

			name  <string>

				Name of the business

		Optional fields:

			type  <string>

				The industry / sector that the business serves and/or operates within (financial services, health/wellness, etc.)

*Example request(s)*

	POST /business
	{
		"owner_id":123,
		"name":"Later Gator LLC",
		"type":"Pest control"
	}

*Response format*

	Success:

		HTTP/1.1 201 Created
		Content-Type: application/json

		{
			"business": {
				"ID": 456,
				"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"DeletedAt": null,
				"owner_id": 123,
				"name": "Later Gator LLC"
			}
		}

	Failure:

		-- Case = Bad request body
		HTTP/1.1 400 Bad Request
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
func (app *Application) CreateBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&business); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	returnRecords, err := business.Create(app.AppDB)
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
		returnRecords)
}

/*
*Description*

func GetBusiness

Get Business record from the database by ID.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/business/{id}

	Body:

		None

*Example request(s)*

	GET /business/456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": null,
			"owner_id": 123,
			"name": "Later Gator LLC"
		}

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Bad Request
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Case = ID does not exist in database
		HTTP/1.1 404 Resource Not Found
		Content-Type: application/json

		{
			"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) GetBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	businessID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnRecords, err := business.Get(app.AppDB, businessID)
	returnedBusiness := returnRecords["business"]
	if err != nil {
		var errorMessage string = fmt.Sprintf("Business ID (%d) does not exist in the database.  [%s]", businessID, err)

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
		returnedBusiness)
}

/*
*Description*

func UpdateBusiness

Updates the Business record associated with the specified business ID in the database.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:   PUT

	Route:  /business/{id}

	Body:
		Format: JSON

		Required fields:

			N/A  --  At least one field should be present in the request body, but no fields are specifically required to be present in the request body.

		Optional fields:

			owner_id  <uint>

				ID of the User record that created/owns the business record

			name  <string>

				Name of the business

*Example request(s)*

	PUT /business/456
	{
		"name": "Sooner Gator, Inc."
	}

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2022-07-11T01:23:45.6789012-14:25",
			"DeletedAt": null,
			"owner_id": 123,
			"name": "Sooner Gator, Inc."
		}

	Failure:
		-- Case = Bad request body or missing/misformatted ID in request URL
		HTTP/1.1 400 Bad Request
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
func (app *Application) UpdateBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	businessID, err := utils.ParseRequestID(request)

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
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	returnRecords, err := business.Update(app.AppDB, businessID, updates)
	updatedBusiness := returnRecords["business"]
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
		updatedBusiness)
}

/*
*Description*

func DeleteBusiness

Delete a Business record from the database by business ID if the ID exists in the database.

Deleted Business record is returned in the response body if the operation is sucessful.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	DELETE

	Route:	/business/{id}

	Body:

		None

*Example request(s)*

	DELETE /business/456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": "2022-07-11T01:23:45.6789012-14:25",
			"owner_id": 123,
			"name": "Later Gator LLC"
		}

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Bad Request
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
func (app *Application) DeleteBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	businessID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnRecords, err := business.Delete(app.AppDB, businessID)
	deletedBusiness := returnRecords["business"]
	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		deletedBusiness)
}

/*
*Description*

func GetBusinesses

Get a list of all Business records in the database.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/businesses

	Body:

		None

*Example request(s)*

	GET /businesses

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		[
			{
				"ID": 727,
				"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"DeletedAt": null,
				"owner_id": 123,
				"name": "Later Gator LLC"
			},
			{
				"ID": 813,
				"CreatedAt": "2022-07-10T14:32:13.1589417-05:00",
				"UpdatedAt": "2022-11-23T05:41:03.4507451-05:00",
				"DeletedAt": null,
				"owner_id": 420,
				"name": "Abraham's Resort & Spa"
			},
			...
		]

	Failure:

		HTTP/1.1 500 InternalServerError
		Content-Type: application/json

		{
			"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) GetBusinesses(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	var businesses []models.Business

	businesses, err := business.GetAll(app.AppDB)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		log.Printf("ERROR:  %s", err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		businesses)
}

/*
*Description*

func GetBusinesses

Get a list of all Services in the database that were created by the specified Business.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/business/{id}/services

	Body:

		None

*Example request(s)*

	GET /business/42/services

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json


		[
			{
				"ID": 11,
				"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"DeletedAt": null,
				"business_id": 42,
				"name":"Yoga class",
				"desc":"30 minute beginner yoga class",
				"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
				"length":30,
				"capacity":20,
				"price":2000,
				"cancel_fee":0,
				"appt_ct":0,
				"is_full":false
			},
			{
				"ID": 83,
				"CreatedAt": "2022-07-10T14:32:13.1589417-05:00",
				"UpdatedAt": "2022-11-23T05:41:03.4507451-05:00",
				"DeletedAt": null,
				"business_id": 42,
				"name":"Caligraphy lessons",
				"desc":"60 minute instructor-led course on caligraphy.",
				"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
				"length":60,
				"capacity":10,
				"price":10000,
				"cancel_fee":2000,
				"appt_ct":0,
				"is_full":false
			},
			...
		]

	Failure:
		-- Case = Bad request body
		HTTP/1.1 400 Bad Request
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Other errors
		HTTP/1.1 500 InternalServerError
		Content-Type: application/json

		{
			"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) GetBusinessServices(writer http.ResponseWriter, request *http.Request) {
	businessID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	service := models.Service{}
	var services []models.Service
	var businessIDJsonKey string = "business_id"

	services, err = service.GetRecordsBySecondaryID(app.AppDB, businessIDJsonKey, businessID)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		log.Printf("ERROR:  %s", err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		services)
}

/*
*Description*

func GetUserServiceAppointments

Get a list of all the Appointments (and the associated Service for each Appointment record) for the specified User.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	GET

	Route:	/user/{id}/service-appointments

	Body:

		None

*Example request(s)*

	GET /user/42/service-appointments

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		[
			{
				"appointment": {
					"ID": 123,
					"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
					"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
					"DeletedAt": null,
					"service_id":11,
					"user_id":42,
					"cancel_date_time":null,
					"active":true
				},
				"service": {
					"ID": 11,
					"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
					"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
					"DeletedAt": null,
					"business_id": 66,
					"name":"Yoga class",
					"desc":"30 minute beginner yoga class",
					"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
					"length":30,
					"capacity":20,
					"price":2000,
					"cancel_fee":0,
					"appt_ct":0,
					"is_full":false
				}
			},
			{
				"appointment": {
					"ID": 456,
					"CreatedAt": "2022-07-10T14:32:13.1589417-05:00",
					"UpdatedAt": "2022-11-23T05:41:03.4507451-05:00",
					"DeletedAt": null,
					"service_id":83,
					"user_id":42,
					"cancel_date_time":null,
					"active":true
				},
				"service": {
					"ID": 83,
					"CreatedAt": "2022-07-10T14:32:13.1589417-05:00",
					"UpdatedAt": "2022-11-23T05:41:03.4507451-05:00",
					"DeletedAt": null,
					"business_id": 91,
					"name":"Caligraphy lessons",
					"desc":"60 minute instructor-led course on caligraphy.",
					"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
					"length":60,
					"capacity":10,
					"price":10000,
					"cancel_fee":2000,
					"appt_ct":0,
					"is_full":false
				}
			},
			...
		]

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Bad Request
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Other errors
		HTTP/1.1 500 Internal Server Error
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) GetBusinessServiceAppointments(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	businessID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	var businessSvcAppts []map[string]interface{}
	businessSvcAppts, err = business.GetServiceAppointments(app.AppDB, businessID, true)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		log.Printf("ERROR:  %s", err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		businessSvcAppts)
}
