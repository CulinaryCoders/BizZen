package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/utils"
)

// TODO:  Add comment documentation (func CreateBusiness)
func (dbHandler *DatabaseHandler) CreateBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&business); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	if err := dbHandler.DB.Create(&business).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusCreated,
		business)
}

// TODO:  Add comment documentation (func CreateBusiness)
func (dbHandler *DatabaseHandler) GetBusiness(writer http.ResponseWriter, request *http.Request) {
}

// TODO:  Add comment documentation (func CreateBusiness)
func (dbHandler *DatabaseHandler) UpdateBusiness(writer http.ResponseWriter, request *http.Request) {
}

// TODO:  Add comment documentation (func CreateBusiness)
func (dbHandler *DatabaseHandler) DeleteBusiness(writer http.ResponseWriter, request *http.Request) {
}
