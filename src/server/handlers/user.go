package handlers

import (
	"encoding/json"
	"net/http"
	"server/config"
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
//	  "username": "johndoe",
//	  "email": "johndoe@example.com"
//	}
func (dbHandler *DatabaseHandler) CreateUser(writer http.ResponseWriter, request *http.Request) {
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

	if err := dbHandler.DB.Create(&user).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusCreated,
		user.ID)
}

// TODO: Add comment documentation (func Authenticate)
func (dbHandler *DatabaseHandler) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&credentials); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	user, err := dbHandler.checkIfUserExists(credentials.Email, writer, request)
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

	validToken, err := GenerateToken(user.Email, user.AccountType, config.AppConfig.GetSigningKey())
	if err != nil {
		utils.RespondWithError(
			writer,
			http.StatusInternalServerError,
			err.Error())
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		validToken)
}

// TODO: Add comment documentation (func checkIfUserExists)
func (dbHandler *DatabaseHandler) checkIfUserExists(userEmail string, writer http.ResponseWriter, request *http.Request) (*models.User, error) {
	var user models.User

	if err := dbHandler.DB.First(&user, models.User{Email: userEmail}).Error; err != nil {
		utils.RespondWithError(writer, http.StatusNotFound, "User does not exist.")
		return nil, err
	}

	return &user, nil
}

// TODO: Add comment documentation (func GetUser)
func (dbHandler *DatabaseHandler) GetUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := dbHandler.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		user)
}

// TODO: Add comment documentation (func UpdateUser)
func (dbHandler *DatabaseHandler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := dbHandler.checkIfUserExists(userEmail, writer, request)
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

	if err := dbHandler.DB.Save(&user).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		user)

}

// TODO: Add comment documentation (func DeleteUser)
func (dbHandler *DatabaseHandler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := dbHandler.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	if err := dbHandler.DB.Delete(&user).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		user)
}
