package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/utils"
)

/*
*Description*

func CreateUser

Creates a new user account record in the database.

*Parameters*

	writer  <http.ResponseWriter>

	   The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	POST

	Route:	/register

	Body:
		Format:	JSON

		Required fields:

	  		email  <string>

				The email address associated with the new user account

	  		password  <string>

				The password for the new user account

			account_type  <string>

				The user account type

				Permitted values:
					User
					Business
					System

			first_name  <string>

				The user's first name

			last_name  <string>

				The user's last name

		Optional fields:

			business_id  <uint>

				The ID of the Business record associated with the new user account (only applicable for 'Business' account types)

*Example request(s)*

	POST /register
	{
	  "email": "johndoe@example.com",
	  "password": "secretpassword",
	  "account_type": "User",
	  "first_name": "John",
	  "last_name": "Doe"
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
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "User",
		"first_name": "John",
		"last_name": "Doe",
		"business_id": null
		}

	Failure:
		-- Case = Bad request body
		HTTP/1.1 400 Bad Request
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Case = Email address already exists in database
		HTTP/1.1 409 Conflict
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
func (app *Application) CreateUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	business := models.Business{}
	var createdUser models.Model
	var createdBusiness models.Model

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	//  Confirm email address doesn't already exist in the DB before creating User.
	//  Check performed here so that 409 status can be passed more easily in the event of a duplicate email.
	isDuplicateEmail, err := user.EmailExists(app.AppDB)
	if isDuplicateEmail {

		utils.RespondWithError(
			writer,
			http.StatusConflict,
			err.Error())

		return

	} else if err != nil {

		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		return
	}

	//  Standardize all User attribute values
	user.StandardizeFields()

	//  Confirm User has valid AccountType
	if !models.UserAccountTypeIsValid(user.AccountType) {
		var errorMessage string = fmt.Sprintf("Invalid account type specified when creating new User record (account_type = %s). Account type must be 'User', 'Business', or 'System'.", user.AccountType)

		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			errors.New(errorMessage).Error())
	}

	createdRecords, err := user.Create(app.AppDB)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		return
	}
	createdUser = createdRecords["user"]

	if user.AccountType == "Business" {
		returnedRecords, err := business.Get(app.AppDB, *user.BusinessID)
		if err != nil {
			utils.RespondWithError(
				writer,
				http.StatusInternalServerError,
				err.Error())

			return
		}
		createdBusiness = returnedRecords["business"]
	}

	returnRecords := map[string]interface{}{
		"user":     createdUser,
		"business": createdBusiness,
	}

	utils.RespondWithJSON(
		writer,
		http.StatusCreated,
		returnRecords)
}

/*
*Description*

func GetUser

Retrieves a user account record from the database by user ID if the ID exists in the database.

*Parameters*

	writer  <http.ResponseWriter>

	   The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	GET

	Route:	/user/{id}

	Body:

		None

*Example request(s)*

	GET /user/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
		"ID": 123456,
		"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
		"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
		"DeletedAt": null,
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "User",
		"first_name": "John",
		"last_name": "Doe",
		"business_id": null
		}

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Bad Request
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}

		-- Case = User ID does not exist in the database
		HTTP/1.1 404 Not Found
		Content-Type: application/json

		{
		"error":"ERROR MESSAGE TEXT HERE"
		}
*/
func (app *Application) GetUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	userID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnRecords, err := user.Get(app.AppDB, userID)
	returnedUser := returnRecords["user"]
	if err != nil {
		var errorMessage string = fmt.Sprintf("User ID (%d) does not exist in the database.  [%s]", userID, err)

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
		returnedUser)
}

/*
*Description*

func UpdateUser

Updates the user record associated with the specified user ID in the database.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "first_name": "").

*Parameters*

	   writer  <http.ResponseWriter>

		   The HTTP response writer

	   request  <*http.Request>

		   The HTTP request

*Returns*

	None

*Expected request format*

	Type:   PUT

	Route:  /user/{id}

	Body:
		Format: JSON

		Required fields:

			N/A  --  At least one field should be present in the request body, but no fields are specifically required to be present in the request body.

		Optional fields:

			email  <string>

				The email address associated with the new user account

			password  <string>

				The password for the new user account

			account_type  <string>

				The user account type

				Permitted values:
					User
					Business
					System

			first_name  <string>

				The user's first name

			last_name  <string>

				The user's last name

			business_id  <uint>

				The ID of the Business record associated with the new user account (only applicable for 'Business' account types)

*Example request(s)*

	PUT /user/123456
	{
		"first_name":"Luke",
		"last_name":"Skywalker"
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
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "User",
		"first_name": "Luke",
		"last_name": "Skywalker",
		"business_id": null
		}

	Failure:
		-- Case = Bad request body or missing/bad ID in request URL
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
func (app *Application) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	userID, err := utils.ParseRequestID(request)

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

	returnRecords, err := user.Update(app.AppDB, userID, updates)
	updatedUser := returnRecords["user"]
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
		updatedUser)

}

/*
*Description*

func DeleteUser

Delete a user account record from the database by user ID if the ID exists in the database.

Deleted user record is returned in the response body if the operation is sucessful.

*Parameters*

	writer  <http.ResponseWriter>

	   The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	DELETE

	Route:	/user/{id}

	Body:

		None

*Example request(s)*

	DELETE /user/123456

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		{
		"ID": 123456,
		"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
		"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
		"DeletedAt": "2022-07-11T01:23:45.6789012-14:25",
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "User",
		"first_name": "John",
		"last_name": "Doe",
		"business_id": null
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
func (app *Application) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	userID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnRecords, err := user.Delete(app.AppDB, userID)
	deletedUser := returnRecords["user"]
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
		deletedUser)
}

/*
*Description*

func GetUsers

Get a list of all User records in the database.

*Parameters*

	writer  <http.ResponseWriter>

		The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:	GET

	Route:	/users

	Body:

		None

*Example request(s)*

	GET /users

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		[
			{
				"ID": 72,
				"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
				"DeletedAt": null,
				"email": "curb-it@example.com",
				"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
				"account_type": "User",
				"first_name": "Larry",
				"last_name": "David",
				"business_id": null
			},
			{
				"ID": 411,
				"CreatedAt": "2022-07-10T14:32:13.1589417-05:00",
				"UpdatedAt": "2022-11-23T05:41:03.4507451-05:00",
				"DeletedAt": null,
				"email": "bubble.guppies.witch@hotmail.com",
				"password": "qwerQEWR174$8O4$1Qfy31MinvsYvPbOCeIXj2fSxMCh8O4$IT",
				"account_type": "Business",
				"first_name": "Wanda",
				"last_name": "Sykes",
				"business_id": 31
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
func (app *Application) GetUsers(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	var users []models.User

	users, err := user.GetAll(app.AppDB)
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
		users)
}

/*
*Description*

func GetUserEnrolledStatus

Returns 'true' if User has an active appointment for the specified Service ('false' if they don't).

*Parameters*

	writer  <http.ResponseWriter>

	   The HTTP response writer

	request  <*http.Request>

		The HTTP request

*Returns*

	None

*Expected request format*

	Type:  	GET

	Route:	/user/{id}/service/{id}

	Body:

		None

*Example request(s)*

	GET /user/420/service/99

*Response format*

	Success:

		HTTP/1.1 200 OK
		Content-Type: application/json

		Body:
		true/false

	Failure:
		-- Case = Required ID field missing from or incorrectly formatted in request url
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
func (app *Application) GetUserEnrolledStatus(writer http.ResponseWriter, request *http.Request) {
	var serviceIDKey string = "service-id"
	var userIDKey string = "user-id"

	serviceID, serviceIDErr := utils.ParseRequestIDField(request, serviceIDKey)
	if serviceIDErr != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			serviceIDErr.Error())

		return
	}

	userID, userIDErr := utils.ParseRequestIDField(request, userIDKey)
	if userIDErr != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			serviceIDErr.Error())

		return
	}

	var user *models.User
	hasServiceAppointment, err := user.HasServiceAppointment(app.AppDB, userID, serviceID)
	if userIDErr != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())

		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		hasServiceAppointment)
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

	POST /user/42/service-appointments

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
func (app *Application) GetUserServiceAppointments(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	userID, err := utils.ParseRequestID(request)

	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	var userSvcAppts []map[string]interface{}
	userSvcAppts, err = user.GetServiceAppointments(app.AppDB, userID, true)
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
		userSvcAppts)
}
