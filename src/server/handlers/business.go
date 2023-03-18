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

Creates a new business and office record in the database and returns them as JSON in the response body.

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
				"main_office_id": 789,
				"name": "Later Gator LLC",
				"type": "Pest control"
			},
			"office": {
				"ID": 789,
				"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"DeletedAt": null,
				"business_id": 456,
				"contact_info_id": null,
				"manager_id": 123,
				"operating_hours_id": null,
				"name": "Later Gator LLC - Main Office"
			}
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
			"main_office_id": 789,
			"name": "Later Gator LLC",
			"type": "Pest control"
		}

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Internal Server Error
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
		var errorMessage string = fmt.Sprintf("Business ID (%d) does not exist in the database.\n%s", businessID, err)

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

			main_office_id  <uint>

				ID of the Office record considered to be the main office for the business record

			name  <string>

				Name of the business

			type  <string>

				The industry / sector that the business serves and/or operates within (financial services, health/wellness, etc.)

*Example request(s)*

	PUT /business/456
	{
		"name": "Sooner Gator, Inc.",
		"main_office_id": 999
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
			"main_office_id": 999,
			"name": "Sooner Gator, Inc.",
			"type": "Pest control"
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
			"main_office_id": 789,
			"name": "Later Gator LLC",
			"type": "Pest control"
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

// TODO:  Update response info in docstring with correct fields/values (func CreateOffice)
/*
*Description*

func CreateOffice

Creates a new office record in the database.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  POST

	Route:	/office

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

	POST /office
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
func (app *Application) CreateOffice(writer http.ResponseWriter, request *http.Request) {
	office := models.Office{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&office); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	returnedRecords, err := office.Create(app.AppDB)
	createdOffice := returnedRecords["office"]
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
		createdOffice)
}

// TODO:  Update response info in docstring with correct fields/values (func GetOffice)
/*
*Description*

func GetOffice

Get office record from the database by ID.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/office/{id}

	Body:

		None

*Example request(s)*

	GET /office/123456

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
func (app *Application) GetOffice(writer http.ResponseWriter, request *http.Request) {
	office := models.Office{}
	officeID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnRecords, err := office.Get(app.AppDB, officeID)
	returnedOffice := returnRecords["office"]
	if err != nil {
		var errorMessage string = fmt.Sprintf("Office ID (%d) does not exist in the database.\n%s", officeID, err)

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
		returnedOffice)
}

// TODO:  Update response info in docstring with correct fields/values (func UpdateOffice)
/*
*Description*

func UpdateOffice

Updates the office record associated with the specified office ID in the database.

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

	Route:  /office/{id}

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

	PUT /office/123456
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
func (app *Application) UpdateOffice(writer http.ResponseWriter, request *http.Request) {
	office := models.Office{}
	officeID, err := utils.ParseRequestID(request)

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

	returnedRecords, err := office.Update(app.AppDB, officeID, updates)
	updatedOffice := returnedRecords["office"]
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
		updatedOffice)
}

// TODO:  Update response info in docstring with correct fields/values (func DeleteOffice)
/*
*Description*

func DeleteOffice

Delete an office record from the database by office ID if the ID exists in the database.

Deleted office record is returned in the response body if the operation is sucessful.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	DELETE

	Route:	/office/{id}

	Body:

		None

*Example request(s)*

	DELETE /office/123456

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
func (app *Application) DeleteOffice(writer http.ResponseWriter, request *http.Request) {
	office := models.Office{}
	officeID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedRecords, err := office.Delete(app.AppDB, officeID)
	deletedOffice := returnedRecords["office"]
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
		deletedOffice)

}
