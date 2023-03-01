package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/utils"
	_ "time"

	"github.com/gorilla/mux"
)

// TODO:  Add documentation (func CreateAddress)
func (app *Application) CreateAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&address); err != nil {
		utils.RespondWithError(
			writer,
			http.StatusBadRequest,
			err.Error())

		return
	}

	defer request.Body.Close()

	createdAddress, err := address.CreateAddress(app.AppDB)
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
		createdAddress)
}

// TODO:  Add comment documentation (func GetAddress)
func (app *Application) GetAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}
	addressID := mux.Vars(request)["id"]

	returnedAddress, err := address.GetAddress(app.AppDB, addressID)
	if err != nil {
		var errorMessage string = fmt.Sprintf("Address ID (%s) does not exist in the database.\n%s", addressID, err)

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
		returnedAddress)
}

// TODO:  Add comment documentation (func UpdateAddress)
func (app *Application) UpdateAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}
	addressID := mux.Vars(request)["id"]

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

	updatedAddress, err := address.UpdateAddress(app.AppDB, addressID, updates)
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
		updatedAddress)
}

// TODO:  Add comment documentation (func DeleteAddress)
func (app *Application) DeleteAddress(writer http.ResponseWriter, request *http.Request) {
	address := models.Address{}
	addressID := mux.Vars(request)["id"]

	deletedAddress, err := address.DeleteAddress(app.AppDB, addressID)
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
		deletedAddress)

}
