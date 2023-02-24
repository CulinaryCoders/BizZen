package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/utils"
)

// TODO:  Add comment documentation (func CreateBusiness)
func (h *Handler) CreateBusiness(writer http.ResponseWriter, request *http.Request) {
	business := models.Business{}

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&business); err != nil {
		utils.RespondWithError(writer, http.StatusBadRequest, err.Error())
		return
	}

	defer request.Body.Close()

	if err := h.DB.Create(&business).Error; err != nil {
		utils.RespondWithError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(
		writer,
		http.StatusCreated,
		business)
}

// TODO:  Add comment documentation (func CreateBusiness)
func (h *Handler) GetBusiness(writer http.ResponseWriter, request *http.Request) {
}

// TODO:  Add comment documentation (func CreateBusiness)
func (h *Handler) UpdateBusiness(writer http.ResponseWriter, request *http.Request) {
}

// TODO:  Add comment documentation (func CreateBusiness)
func (h *Handler) DeleteBusiness(writer http.ResponseWriter, request *http.Request) {
}
