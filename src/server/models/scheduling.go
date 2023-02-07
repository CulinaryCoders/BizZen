package models

import (
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Appointment model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type Appointment struct {
	gorm.Model
	ID                uint `gorm:"primaryKey;serial"`
	ClientID          uint
	ServiceOfferingID uint
	ScheduledDatetime time.Time
	CreatedDatetime   time.Time
	Active            bool
	Approved          bool
}

// TODO: Update time columns type / formatting to ensure behavior/values are expected
type TimeSlots struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey;serial"`
	TimeSlot time.Time `gorm:"not null"`
}

// TODO: Add foreign key logic to StaffShifts model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type StaffShifts struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;serial"`
	StaffID   uint
	DayOfWeek string
	StartTime time.Time
	EndTime   time.Time
}

// TODO: Add foreign key logic to ResourceAvailability model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type ResourceAvailability struct {
	gorm.Model
	ID         uint `gorm:"primaryKey;serial"`
	ResourceID uint
	DayOfWeek  string
	StartTime  time.Time
	EndTime    time.Time
}
