package handlers

import (
	"gorm.io/gorm"
)

// TODO: Add comment documentation (type Handler)
type Handler struct {
	DB *gorm.DB
}

// TODO: Add comment documentation (func NewHandler)
func NewHandler(db *gorm.DB) Handler {
	return Handler{db}
}
