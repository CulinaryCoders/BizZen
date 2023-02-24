package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/utils"

	"github.com/gorilla/mux"
)

// TODO: Add comment documentation (type Credentials)
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser is an HTTP handler that creates a new user.
//
// This handler expects a POST request with a JSON body containing the following fields:
//   - "username" (string): the username of the new user
//   - "email" (string): the email address of the new user
//   - "password" (string): the password for the new user
//
// If the user is successfully created, this handler returns a JSON response with the following fields:
//   - "id" (string): the unique ID of the new user
//   - "username" (string): the username of the new user
//   - "email" (string): the email address of the new user
//
// If there is an error creating the user (e.g. if the username is already taken), this handler returns a JSON response with the following fields:
//   - "error" (string): a message describing the error that occurred
//
// Example usage:
//
//	POST /users
//	{
//	  "username": "johndoe",
//	  "email": "johndoe@example.com",
//	  "password": "secretpassword"
//	}
//
// Response:
//
//	HTTP/1.1 201 Created
//	Content-Type: application/json
//
//	{
//	  "id": "123456",
//	}
func (h *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&user); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	if err := user.HashPassword(user.Password); err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusCreated,
		user.ID)
}

// AuthenticateUser is an HTTP handler that authenticates a user.
//
// This handler expects a POST request with a JSON body containing the following fields:
//   - "email" (string): the username of the user to authenticate
//   - "password" (string): the password of the user to authenticate
//
// If the user is successfully authenticated, this handler returns a JSON response with the following fields:
//   - "token" (string): a JWT token that can be used to authorize future requests
//
// If there is an error authenticating the user (e.g. if the username or password is incorrect), this handler returns a JSON response with the following fields:
//   - "error" (string): a message describing the error that occurred
//
// Example usage:
//
//	POST /login
//	{
//	  "email": "johndoe@example.com",
//	  "password": "secretpassword"
//	}
//
// Response:
//
//	HTTP/1.1 200 OK
//	Content-Type: application/json
//
//	{
//	  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
//	}
func (h *Handler) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&credentials); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	user, err := h.checkIfUserExists(credentials.Email, writer, request)
	if err != nil {
		return
	}

	if err := user.CheckPassword(credentials.Password); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			"Incorrect password.")

		return
	}
	session, _ := h.Store.Get(request, "sessionID")
	session.Values["authenticated"] = true
	session.Save(request, writer)
	//validToken, err := GenerateToken(user.Email, user.AccountType, config.AppConfig.GetSigningKey())
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

// checkIfUserExists is a helper function that checks if a user with the specified email exists in the database.
//
// Parameters:
//   - email (string): the email of the user to check.
//   - writer (http.ResponseWriter): an HTTP response writer for writing the response.
//   - request (*http.Request): an HTTP request object containing the user email in the URL path.
//
// Returns:
//   - *User: a pointer to the User object if the user exists, nil otherwise.
//   - error: a 404 Not Found error if there was a problem checking if the user exists.
//
// Example usage:
//
//   func GetUser(writer http.ResponseWriter, request *http.Request) {
//     	email := mux.Vars(request)["email"]
//
//		user, err := h.checkIfUserExists(userEmail, writer, request)
// 		if err != nil {
// 			return
// 		}

//		utils.RespondWithJSON(
//			writer,
//			http.StatusOK,
//			user)
//	  }
func (h *Handler) checkIfUserExists(userEmail string, writer http.ResponseWriter, request *http.Request) (*models.User, error) {
	var user models.User

	if err := h.DB.First(&user, models.User{Email: userEmail}).Error; err != nil {
		utils.RespondWithError(writer, http.StatusNotFound, "User does not exist.")
		return nil, err
	}

	return &user, nil
}

// GetUser is an HTTP handler that creates a new user.
//
// This handler expects a GET request with a URL path that includes the Email of the user to retrieve:
//   - GET /users/{email}
//
// Response:
//   - If the user is successfully found, the handler function responds with a JSON-encoded User object.
//
// If there is an error getting the user (e.g. if the email does not exist), this handler returns a JSON response with the following fields:
//   - "error" (string): a message describing the error that occurred
func (h *Handler) GetUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := h.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		user)
}

// UpdateUser is an HTTP handler function that updates a user's information in the database and responds with a JSON-encoded User object.
//
// Parameters:
//   - writer (http.ResponseWriter): an HTTP response writer for writing the response.
//   - request (*http.Request): an HTTP request object containing the user email in the URL path and the updated user data in the request body.
//
// HTTP Request:
//   The handler function expects a PUT or POST request with a URL path that includes the email of the user to update:
//
//      PUT /users/{email}
//      POST /users/{email}
//
//   The {email} path parameter should be replaced with the email of the user to update.
//
//   The request body should contain a JSON object with the updated user data. For example:
//
//      {
//        "first_name": "New Name",
//        "email": "new-email@example.com"
//      }
//
// Returns:
//   - none
//
// Response:
//   The handler function responds with a JSON-encoded User object representing the updated user. If the user is not found in the database, the function responds with a 404 Not Found error. If the request body is invalid or the update fails for some other reason, the function responds with a 400 Bad Request error or a 500 Internal Server error.
//

func (h *Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := h.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	if err := h.DB.Save(&user).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		user)

}

// DeleteUserByEmail is an HTTP handler function that deletes a user from the database by email and responds with a JSON-encoded success message.
//
// Parameters:
//   - writer (http.ResponseWriter): an HTTP response writer for writing the response.
//   - request (*http.Request): an HTTP request object containing the user email in the URL path.
//
// HTTP Request:
//
//	The handler function expects a DELETE request with a URL path that includes the email of the user to delete:
//
//	   DELETE /users/email/{email}
//
//	The {email} path parameter should be replaced with the email of the user to delete.
//
// Returns:
//   - none
//
// Response:
//
//	The handler function responds with a JSON-encoded success message indicating that the user has been successfully deleted. If the user is not found in the database, the function responds with a 404 Not Found error. If the delete operation fails for some other reason, the function responds with a 500 Internal Server Error.
func (h *Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := h.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		user)
}
