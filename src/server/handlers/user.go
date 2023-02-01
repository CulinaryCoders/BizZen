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

func (h handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := user.HashPassword(user.Password); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, user)
}

func (h handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&credentials); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	user, err := h.checkIfUserExists(credentials.Email, w, r)
	if err != nil {
		return
	}

	if err := user.CheckPassword(credentials.Password); err != nil {
		respondError(w, http.StatusBadRequest, "Incorrect password.")
		return
	}

	validToken, err := controllers.GenerateToken(user.Email, user.Role)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	respondJSON(w, http.StatusOK, validToken)
}

func (h handler) checkIfUserExists(email string, w http.ResponseWriter, r *http.Request) (*models.User, error) {
	var user models.User

	if err := h.DB.First(&user, models.User{Email: email}).Error; err != nil {
		respondError(w, http.StatusNotFound, "User does not exist.")
		return nil, err
	}
	return &user, nil
}

func (h handler) FindUser(w http.ResponseWriter, r *http.Request) {
	userEmail := mux.Vars(r)["email"]

	user, err := h.checkIfUserExists(userEmail, w, r)
	if err != nil {
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func (h handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["email"]

	user, err := h.checkIfUserExists(userId, w, r)
	if err != nil {
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := h.DB.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)

}

func (h handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userEmail := mux.Vars(r)["email"]

	user, err := h.checkIfUserExists(userEmail, w, r)
	if err != nil {
		return
	}

	if err := h.DB.Delete(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, user)
}
