package handlers

import (
	"gorm.io/gorm"
)

// TODO: Add comment documentation (type Handler)
type DatabaseHandler struct {
	DB *gorm.DB
}

// TODO: Add comment documentation (func NewHandler)
func NewDatabaseHandler(db *gorm.DB) *DatabaseHandler {
	return &DatabaseHandler{db}
}
