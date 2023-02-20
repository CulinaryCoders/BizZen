package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/utils"
)

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
