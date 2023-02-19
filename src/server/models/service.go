package models

import (
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Service model
type Service struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;serial"`
	OfficeID    uint
	Name        string
	Description string
}

// TODO: Add foreign key logic to ServiceOffering model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type ServiceOffering struct {
	gorm.Model
	ID                     uint `gorm:"primaryKey;serial"`
	ServiceID              uint
	StaffID                uint
	ResourceID             uint
	StartDate              time.Time
	EndDate                time.Time
	BookingLength          uint
	Price                  uint
	CancellationFee        uint
	MaxConsecutiveBookings uint
	MinCancellationNotice  uint
	MinTimeBetweenClients  uint
}
