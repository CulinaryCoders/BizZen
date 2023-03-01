package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/models"
	"server/utils"
	"strconv"

	"github.com/gorilla/mux"
)

/*
The Credentials struct defines the format for user login credentials. It contains two fields: Email and Password. The Email field is a string that represents the user's email address, while the Password field is a string that represents the user's password.
*/
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
Helper function to parse the user's id from the request
*/
func parseRequestID(request *http.Request) (uint64, error) {
	userId := mux.Vars(request)["id"]
	convertedToUint64, err := strconv.ParseUint(userId, 10, 64)
	fmt.Print(convertedToUint64)
	return convertedToUint64, err
}

/*
CreateUser is an HTTP handler that creates a new user.

This handler expects a POST request with a JSON body containing the following fields:
  - "username" (string): the username of the new user
  - "email" (string): the email address of the new user
  - "password" (string): the password for the new user

If the user is successfully created, this handler returns a JSON response with the following field:
  - "id" (string): the unique ID of the new user

If there is an error creating the user (e.g. if the username is already taken), this handler returns a JSON response with the following fields:
  - "error" (string): a message describing the error that occurred

Example usage:

	POST /users
	{
	  "username": "johndoe",
	  "email": "johndoe@example.com",
	  "password": "secretpassword"
	}

Response:

	HTTP/1.1 201 Created
	Content-Type: application/json

	{
	  "id": "123456",
	}
*/
func (env *Env) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	user, err := decodeUser(request)
	if err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	insertedID, err := env.users.CreateUser(user)
	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(writer, http.StatusCreated, insertedID)
}

/*
Helper function to unmarshal the request body into the User model
*/
func decodeUser(request *http.Request) (*models.User, error) {
	user := &models.User{}
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(user); err != nil {
		return nil, err
	}
	defer request.Body.Close()
	return user, nil
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
func (env *Env) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&credentials); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	user, err := env.users.FindUserByEmail(credentials.Email)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			"User not found.")
		return
	}

	if err := user.CheckPassword(credentials.Password); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			"Incorrect password.")

		return
	}
	/*
		PLEASE DO NOT REMOVE
		TODO: Implement this later...
		session, _ := env.Store.Get(request, "sessionID")
		session.Values["authenticated"] = true
		session.Save(request, writer)
		//validToken, err := GenerateToken(user.Email, user.AccountType, config.AppConfig.GetSigningKey())
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
func (env *Env) GetUser(writer http.ResponseWriter, request *http.Request) {
	userId, err := parseRequestID(request)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())
	}

	user, err := env.users.FindUser(userId)
	if err != nil {
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		user)
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
func (env *Env) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	userId, err := parseRequestID(request)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())
		return
	}

	updatedUser, err := decodeUser(request)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())
		return
	}

	updateID, err := env.users.UpdateUser(userId, updatedUser)

	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		updateID)

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
func (env *Env) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userId, err := parseRequestID(request)
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())
		return
	}

	deleteSuccess, err := env.users.DeleteUser(userId)
	if err != nil || !deleteSuccess {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		"User deleted")
}
