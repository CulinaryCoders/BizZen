package models

import (
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Business model
type Business struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;serial"`
	OwnerID      uint
	AddressID    uint
	PhoneNumber1 string
	PhoneNumber2 string
	FaxNumber    string
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
