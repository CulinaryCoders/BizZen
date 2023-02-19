package models

import (
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Business model
type Business struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey;serial" json:"id"`
	OwnerID      uint   `json:"owner"`
	MainOfficeID uint   `json:"main_office"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

// TODO: Add foreign key logic to Office model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type Office struct {
	gorm.Model
	ID            uint `gorm:"primaryKey;serial"`
	BusinessID    uint
	ContactInfoID uint
	ManagerID     uint
	Name          string `gorm:"not null"`
	OpeningTime   time.Time
	ClosingTime   time.Time
}
