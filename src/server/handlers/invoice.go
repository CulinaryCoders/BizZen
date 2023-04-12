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

func CreateInvoice

Creates a new invoice record in the database.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  POST

	Route:	/invoice

	Body:
		Format: JSON

		Required fields:

			appointment_id  <uint>

				ID of Appointment record Invoice is associated with

			original_balance  <int>

				Total original balance of the invoice (in cents)

		Optional fields:

			remaining_balance  <int>

				Remaining balance of the invoice (in cents)

			status  <string>

				Status of the invoice.

				Permitted values (automatically set based on remaining balance value):
					Unpaid
					Paid
					Overpaid

*Example request(s)*

	POST /invoice
	{
		"appointment_id":123,
		"original_balance":5000
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
			"appointment_id":123,
			"original_balance":5000,
			"remaining_balance":5000,
			"status":"Unpaid"
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
func (app *Application) CreateInvoice(writer http.ResponseWriter, request *http.Request) {
	invoice := models.Invoice{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&invoice); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	returnedRecords, err := invoice.Create(app.AppDB)
	createdInvoice := returnedRecords["invoice"]
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
		createdInvoice)
}

/*
*Description*

func GetInvoice

Get invoice record from the database by ID.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/invoice/{id}

	Body:

		None

*Example request(s)*

	GET /invoice/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"DeletedAt": null,
			"appointment_id":123,
			"original_balance":5000,
			"remaining_balance":5000,
			"status":"Unpaid"
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
func (app *Application) GetInvoice(writer http.ResponseWriter, request *http.Request) {
	invoice := models.Invoice{}
	invoiceID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedInvoice, err := invoice.Get(app.AppDB, invoiceID)
	if err != nil {
		var errorMessage string = fmt.Sprintf("Invoice ID (%d) does not exist in the database.  [%s]", invoiceID, err)

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
		returnedInvoice)
}

/*
*Description*

func UpdateInvoice

Updates the invoice record associated with the specified invoice ID in the database.

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

	Route:  /invoice/{id}

	Body:
		Format: JSON

		Required fields:

			N/A  --  At least one field should be present in the request body, but no fields are specifically required to be present in the request body.

		Optional fields:

			appointment_id  <uint>

				ID of Appointment record Invoice is associated with

			original_balance  <int>

				Total original balance of the invoice (in cents)

			remaining_balance  <int>

				Remaining balance of the invoice (in cents)

			status  <string>

				Status of the invoice (Unpaid, Paid, Overpaid)

*Example request(s)*

	PUT /invoice/123456
	{
		"remaining_balance":0
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
			"appointment_id":123,
			"original_balance":5000,
			"remaining_balance":0,
			"status":"Paid"
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
func (app *Application) UpdateInvoice(writer http.ResponseWriter, request *http.Request) {
	invoice := models.Invoice{}
	invoiceID, err := utils.ParseRequestID(request)

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

	returnedRecords, err := invoice.Update(app.AppDB, invoiceID, updates)
	updatedInvoice := returnedRecords["invoice"]
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
		updatedInvoice)
}

/*
*Description*

func DeleteInvoice

Delete an invoice record from the database by invoice ID if the ID exists in the database.

Deleted invoice record is returned in the response body if the operation is sucessful.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	DELETE

	Route:	/invoice/{id}

	Body:

		None

*Example request(s)*

	DELETE /invoice/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
			"ID": 123456,
			"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
			"UpdatedAt": "2020-02-13T04:20:12.6789012-05:00",
			"DeletedAt": "2022-06-31T04:20:12.6789012-05:00",,
			"appointment_id":123,
			"original_balance":5000,
			"remaining_balance":0,
			"status":"Paid"
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
func (app *Application) DeleteInvoice(writer http.ResponseWriter, request *http.Request) {
	invoice := models.Invoice{}
	invoiceID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnedRecords, err := invoice.Delete(app.AppDB, invoiceID)
	deletedInvoice := returnedRecords["invoice"]
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
		deletedInvoice)

}
