package handlers

import (
	"encoding/json"
	"net/http"
	"server/controllers"
	"server/models"

	"github.com/gorilla/mux"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (db Handler) RegisterUser(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(writer, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	if err := user.HashPassword(user.Password); err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(writer, http.StatusCreated, user)
}

func (db Handler) LogIn(writer http.ResponseWriter, request *http.Request) {
	var credentials Credentials

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&credentials); err != nil {
		respondError(writer, http.StatusBadRequest, err.Error())
		return
	}
	defer request.Body.Close()

	user, err := db.checkIfUserExists(credentials.Email, writer, request)
	if err != nil {
		return
	}

	if err := user.CheckPassword(credentials.Password); err != nil {
		respondError(writer, http.StatusBadRequest, "Incorrect password.")
		return
	}

	validToken, err := controllers.GenerateToken(user.Email, user.Role)
	if err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
	}

	respondJSON(writer, http.StatusOK, validToken)
}

func (db Handler) checkIfUserExists(email string, writer http.ResponseWriter, request *http.Request) (*models.User, error) {
	var user models.User

	if err := db.DB.First(&user, models.User{Email: email}).Error; err != nil {
		respondError(writer, http.StatusNotFound, "User does not exist.")
		return nil, err
	}
	return &user, nil
}

func (db Handler) FindUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := db.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}
	respondJSON(writer, http.StatusOK, user)
}

func (db Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	userId := mux.Vars(request)["email"]

	user, err := db.checkIfUserExists(userId, writer, request)
	if err != nil {
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		respondError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	if err := db.DB.Save(&user).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(writer, http.StatusOK, user)

}

func (db Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {
	userEmail := mux.Vars(request)["email"]

	user, err := db.checkIfUserExists(userEmail, writer, request)
	if err != nil {
		return
	}

	if err := db.DB.Delete(&user).Error; err != nil {
		respondError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(writer, http.StatusOK, user)
}
