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

// TODO:  Add comment documentation (func CreateBusiness)
func (app *Application) CreateBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&business); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	createdBusiness, err := business.CreateBusiness(app.AppDB)
	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusCreated,
		createdBusiness)
}

// TODO:  Add comment documentation (func GetBusiness)
func (app *Application) GetBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	businessID := mux.Vars(request)["id"]

	returnedBusiness, err := business.GetBusiness(app.AppDB, businessID)
	if err != nil {
		var errorMessage string = fmt.Sprintf("Business ID (%s) does not exist in the database.\n%s", businessID, err)

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
		returnedBusiness)
}

// TODO:  Add comment documentation (func UpdateBusiness)
func (app *Application) UpdateBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	businessID := mux.Vars(request)["id"]

	var updates map[string]interface{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&updates); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	updatedBusiness, err := business.UpdateBusiness(app.AppDB, businessID, updates)
	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		updatedBusiness)
}

// TODO:  Add comment documentation (func DeleteBusiness)
func (app *Application) DeleteBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}
	businessID := mux.Vars(request)["id"]

	deletedBusiness, err := business.DeleteBusiness(app.AppDB, businessID)
	if err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusOK,
		deletedBusiness)
}
