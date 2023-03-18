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

			username  <string>

				The username for the new user account

	  		email  <string>

				The email address associated with the new user account

	  		password  <string>

				The password for the new user account

			account_type  <string>

				The user account type

				Permitted values:
					Individual
					Business
					System

			first_name  <string>

				The user's first name

			last_name  <string>

				The user's last name

		Optional fields:

			contact_info_id  <uint>

				The ID of the ContactInfo record associated with the new user account

			business_id  <uint>

				The ID of the Business record associated with the new user account (only applicable for 'Business' account types)

			permissions_id  <uint>

				The ID of the UserPermissions record associated with the new user account

			user_pref_id  <uint>

				The ID of the UserPrefences record associated with the new user account

			profile_pic_id  <uint>

				The ID of the ProfilePic record associated with the new user account

*Example request(s)*

	POST /register
	{
	  "username": "johndoe",
	  "email": "johndoe@example.com",
	  "password": "secretpassword",
	  "account_type": "Individual",
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
		"username": "johndoe",
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "Individual",
		"first_name": "John",
		"last_name": "Doe",
		"contact_info_id": 45,
		"business_id": null,
		"permissions_id": 123,
		"user_pref_id": 88,
		"profile_pic_id": 79
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
func (app *Application) CreateUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	if err := user.HashPassword(user.Password); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	returnRecords, err := user.Create(app.AppDB)
	createdUser := returnRecords["user"]
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
		createdUser)
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
		"username": "johndoe",
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "Individual",
		"first_name": "John",
		"last_name": "Doe",
		"contact_info_id": 45,
		"business_id": null,
		"permissions_id": 123,
		"user_pref_id": 88,
		"profile_pic_id": 79
		}

	Failure:
		-- Case = ID missing from or incorrectly formatted in request url
		HTTP/1.1 400 Internal Server Error
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
		var errorMessage string = fmt.Sprintf("User ID (%d) does not exist in the database.\n%s", userID, err)

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

			username  <string>

				The username for the new user account

			email  <string>

				The email address associated with the new user account

			password  <string>

				The password for the new user account

			account_type  <string>

				The user account type

				Permitted values:
					Individual
					Business
					System

			first_name  <string>

				The user's first name

			last_name  <string>

				The user's last name

			contact_info_id  <uint>

				The ID of the ContactInfo record associated with the new user account

			business_id  <uint>

				The ID of the Business record associated with the new user account (only applicable for 'Business' account types)

			permissions_id  <uint>

				The ID of the UserPermissions record associated with the new user account

			user_pref_id  <uint>

				The ID of the UserPrefences record associated with the new user account

			profile_pic_id  <uint>

				The ID of the ProfilePic record associated with the new user account

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
		"username": "johndoe",
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "Individual",
		"first_name": "Luke",
		"last_name": "Skywalker",
		"contact_info_id": 45,
		"business_id": null,
		"permissions_id": 123,
		"user_pref_id": 88,
		"profile_pic_id": 79
		}

	Failure:
		-- Case = Bad request body or missing/bad ID in request URL
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
		"username": "johndoe",
		"email": "johndoe@example.com",
		"password": "$2a$14$ITcK9ZosVTZpx3OeJT8qu.I1Qfy31MinvsYvPbOCeIXj2fSxMCh8O",
		"account_type": "Individual",
		"first_name": "John",
		"last_name": "Doe",
		"contact_info_id": 45,
		"business_id": null,
		"permissions_id": 123,
		"user_pref_id": 88,
		"profile_pic_id": 79
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
