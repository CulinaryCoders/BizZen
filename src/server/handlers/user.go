package handlers

import (
	"encoding/json"
	"net/http"
	"server/controllers"
	"server/models"

	"github.com/gorilla/mux"
)

// TODO: Add comment documentation (type Credentials)
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: Add comment documentation (func RegisterUser)
func (db Handler) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(writer, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	// ? Should error handling and response be handled by the called function instead?
	if err := user.HashPassword(user.Password); err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	// ? Should error handling and response be handled by the called function instead?
	// TODO: Create wrapper function for Handler type/struct to encapsulate "gorm.DB" logic
	if err := db.DB.Create(&user).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(
		writer,
		http.StatusCreated,
		user,
	)
}

// TODO: Add comment documentation (func Authenticate)
func (db Handler) Authenticate(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&credentials); err != nil {
		respondError(writer, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	// ? Duplicative code block for checking if user exists. Should create a new type/struct to store user state and reference that instead.
	// TODO: Consolidate block into separate function and/or store user exists check in a new type/struct.
	user, err := db.checkIfUserExists(credentials.Email, writer, request)
	if err != nil {
		return
	}

	// ? Should error handling and response be handled by the called function instead?
	if err := user.CheckPassword(credentials.Password); err != nil {
		respondError(
			writer,
			http.StatusBadRequest,
			"Incorrect password.",
		)

		return
	}

	// ? Should error handling and response be handled by the called function instead?
	validToken, err := controllers.GenerateToken(user.Email, user.Role)
	if err != nil {
		respondError(
			writer,
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	respondJSON(
		writer,
		http.StatusOK,
		validToken,
	)
}

// TODO: Add comment documentation (func checkIfUserExists)
func (db Handler) checkIfUserExists(userEmail string, writer http.ResponseWriter, request *http.Request) (*models.User, error) {
	var user models.User

	// ? Should error handling and response be handled by the called function instead?
	// TODO: Create wrapper function for Handler type/struct to encapsulate "gorm.DB" logic (First / checkIfUserExists)
	if err := db.DB.First(&user, models.User{Email: userEmail}).Error; err != nil {
		respondError(writer, http.StatusNotFound, "User does not exist.")
		return nil, err
	}

	return &user, nil
}

// TODO: Add comment documentation (func FindUser)
func (db Handler) FindUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	// ? Duplicative code block for checking if user exists. Should create a new type/struct to store user state and reference that instead.
	// TODO: Consolidate block into separate function and/or store user exists check in a new type/struct.
	user, err := db.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	respondJSON(
		writer,
		http.StatusOK,
		user,
	)
}

// TODO: Add comment documentation (func UpdateUser)
func (db Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	// ? Duplicative code block for checking if user exists. Should create a new type/struct to store user state and reference that instead.
	// TODO: Consolidate block into separate function and/or store user exists check in a new type/struct.
	user, err := db.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	// ? Duplicative code block for decoding request body and error checking/response.
	// TODO:  Create new function to consolidate duplicative code (decoding request body / error handling).
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	// ? Should error handling and response be handled by the called function instead?
	// TODO: Create wrapper function for Handler type/struct to encapsulate "gorm.DB" logic (Save / UpdateUser)
	if err := db.DB.Save(&user).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(
		writer,
		http.StatusOK,
		user,
	)

}

// TODO: Add comment documentation (func DeleteUser)
func (db Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	// ? Duplicative code block for checking if user exists. Should create a new type/struct to store user state and reference that instead.
	// TODO: Consolidate block into separate function and/or store user exists check in a new type/struct.
	user, err := db.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	// ? Should error handling and response be handled by the called function instead?
	// TODO: Create wrapper function for Handler type/struct to encapsulate "gorm.DB" logic (DeleteUser)
	if err := db.DB.Delete(&user).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(
		writer,
		http.StatusOK,
		user,
	)
}
