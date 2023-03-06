package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/utils"

	"github.com/gorilla/mux"
)

/*
Credentials struct defines the format for user login credentials. It contains two fields: Email and Password. The Email field is a string that represents the user's email address, while the Password field is a string that represents the user's password.
*/
type Credentials struct {
	Email    string `json:"email"`    // User's email
	Password string `json:"password"` // User's password
}

/*
Creates a new user account record in the database.

# Expected request format

	Type:  POST

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

# Example request(s)

	POST /users
	{
	  "username": "johndoe",
	  "email": "johndoe@example.com",
	  "password": "secretpassword",
	  "account_type": "Individual",
	  "first_name": "John",
	  "last_name": "Doe"
	}

# Response format

	Success:

		HTTP/1.1 201 Created
		Content-Type: application/json

		{
		"ID": "123456",
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
			http.StatusInternalServerError,
			err.Error())

		return
	}

	createdUser, err := user.CreateUser(app.AppDB)
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
Authenticate is an HTTP handler that authenticates a user.

This handler expects a POST request with a JSON body containing the following fields:
- "email" (string): the username of the user to authenticate
- "password" (string): the password of the user to authenticate

If the user is successfully authenticated, this handler returns a JSON response indicating the user has logged in
and sets a session cookie in the browser.

If there is an error authenticating the user (e.g. if the username or password is incorrect), this handler returns a JSON response with the following fields:
- "error" (string): a message describing the error that occurred

Example usage:

	POST /login
	{
	"email": "johndoe@example.com",
	"password": "secretpassword"
	}

Response:

	HTTP/1.1 200 OK
	Content-Type: application/json
	Payload: "User logged in."
*/
func (app *Application) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials
	user := models.User{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&credentials); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	returnedUser, err := user.GetUserByEmail(app.AppDB, credentials.Email)
	if err != nil {
		return
	}

	if err := returnedUser.CheckPassword(credentials.Password); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			"Incorrect password.")

		return
	}
	/*
		PLEASE DO NOT REMOVE
		TODO: Implement authentication logic (func Authenticate)
		session, _ := env.Store.Get(request, "sessionID")
		session.Values["authenticated"] = true
		session.Save(request, writer)
		//validToken, err := utils.GenerateToken(user.Email, user.AccountType, config.AppConfig.GetSigningKey())
	*/
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		"User logged in.")
}

/*
GetUser is an HTTP handler that creates a new user.

This handler expects a GET request with a URL path that includes the Email of the user to retrieve:
  - GET /users/{email}

Response:
  - If the user is successfully found, the handler function responds with a JSON-encoded User object.

If there is an error getting the user (e.g. if the email does not exist), this handler returns a JSON response with the following fields:
  - "error" (string): a message describing the error that occurred
*/
func (app *Application) GetUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	userID := mux.Vars(request)["id"]

	// DO NOT DELETE -- KEPT FOR REFERENCE PURPOSES
	// userID, err := utils.ParseRequestID(request)
	// if err != nil {
	// 	utils.RespondWithError(
	// 		writer,
	// 		http.StatusInternalServerError,
	// 		err.Error())
	// }

	returnedUser, err := user.GetUser(app.AppDB, userID)
	if err != nil {
		var errorMessage string = fmt.Sprintf("User ID (%s) does not exist in the database.\n%s", userID, err)

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
UpdateUser is an HTTP handler function that updates a user's information in the database and responds with a JSON-encoded User object.

Parameters:
  - writer (http.ResponseWriter): an HTTP response writer for writing the response.
  - request (*http.Request): an HTTP request object containing the user email in the URL path and the updated user data in the request body.

HTTP Request:

	The handler function expects a PUT or POST request with a URL path that includes the email of the user to update:

	   PUT /users/{email}
	   POST /users/{email}

	The {email} path parameter should be replaced with the email of the user to update.

	The request body should contain a JSON object with the updated user data. For example:

	   {
	     "first_name": "New Name",
	     "email": "new-email@example.com"
	   }

Returns:
  - none

Response:

	The handler function responds with a JSON-encoded User object representing the updated user. If the user is not found in the database, the function responds with a 404 Not Found error. If the request body is invalid or the update fails for some other reason, the function responds with a 400 Bad Request error or a 500 Internal Server error.
*/
func (app *Application) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	userID := mux.Vars(request)["id"]

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

	updatedUser, err := user.UpdateUser(app.AppDB, userID, updates)
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
DeleteUser is an HTTP handler function that deletes a user from the database by email and responds with a JSON-encoded success message.

Parameters:
  - writer (http.ResponseWriter): an HTTP response writer for writing the response.
  - request (*http.Request): an HTTP request object containing the user email in the URL path.

HTTP Request:

	The handler function expects a DELETE request with a URL path that includes the email of the user to delete:

	   DELETE /users/email/{email}

	The {email} path parameter should be replaced with the email of the user to delete.

Returns:
  - none

Response:

	The handler function responds with a JSON-encoded success message indicating that the user has been successfully deleted. If the user is not found in the database, the function responds with a 404 Not Found error. If the delete operation fails for some other reason, the function responds with a 500 Internal Server Error.
*/
func (app *Application) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}
	userID := mux.Vars(request)["id"]

	deletedUser, err := user.DeleteUser(app.AppDB, userID)
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
