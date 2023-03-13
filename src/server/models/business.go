package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Business model
// GORM model for all Business records in the database
type Business struct {
	gorm.Model
	OwnerID      uint   `gorm:"column:owner_id" json:"owner_id"`             // ID of User account that owns the business record
	MainOfficeID uint   `gorm:"column:main_office_id" json:"main_office_id"` // ID of the main Office record for the business record
	Name         string `gorm:"column:name" json:"name"`                     // Business name
	Type         string `gorm:"column:type" json:"type"`                     // The industry / sector that the business serves and/or operates within (financial services, health/wellness, etc.)
}

// TODO: Add foreign key logic to Office model
// GORM model for all Office records in the database
type Office struct {
	gorm.Model
	BusinessID       uint   `gorm:"not null;column:business_id" json:"business_id"`      // ID of the Business record the Office record is associated with
	ContactInfoID    uint   `gorm:"column:contact_info_id" json:"contact_info_id"`       // ID of the ContactInfo record that stores the contact information for the Office record
	ManagerID        uint   `gorm:"not null;column:manager_id" json:"manager_id"`        // ID of User account that manages this Office record
	OperatingHoursID uint   `gorm:"column:operating_hours_id" json:"operating_hours_id"` // ID of OperatingHours record that stores the operating hours for the Office record
	Name             string `gorm:"not null;column:name" json:"name"`                    // Office name/label
}

// TODO: Test time columns type / formatting to ensure behavior/values are expected
// GORM model for all OperatingHours records in the database
type OperatingHours struct {
	gorm.Model
	// Sunday
	OpenSunday      bool      `gorm:"column:open_sunday" json:"open_sunday"`             // Open on Sundays (boolean)
	SundayOpenTime  time.Time `gorm:"column:sunday_open_time" json:"sunday_open_time"`   // Sunday opening time (24 hour format)
	SundayCloseTime time.Time `gorm:"column:sunday_close_time" json:"sunday_close_time"` // Sunday closing time (24 hour format)
	// Monday
	OpenMonday      bool      `gorm:"column:open_monday" json:"open_monday"`             // Open on Mondays (boolean)
	MondayOpenTime  time.Time `gorm:"column:monday_open_time" json:"monday_open_time"`   // Monday opening time (24 hour format)
	MondayCloseTime time.Time `gorm:"column:monday_close_time" json:"monday_close_time"` // Monday closing time (24 hour format)
	// Tuesday
	OpenTuesday      bool      `gorm:"column:open_tuesday" json:"open_tuesday"`             // Open on Tuesdays (boolean)
	TuesdayOpenTime  time.Time `gorm:"column:tuesday_open_time" json:"tuesday_open_time"`   // Tuesday opening time (24 hour format)
	TuesdayCloseTime time.Time `gorm:"column:tuesday_close_time" json:"tuesday_close_time"` // Tuesday closing time (24 hour format)
	// Wednesday
	OpenWednesday      bool      `gorm:"column:open_wednesday" json:"open_wednesday"`             // Open on Wednesdays (boolean)
	WednesdayOpenTime  time.Time `gorm:"column:wednesday_open_time" json:"wednesday_open_time"`   // Wednesday opening time (24 hour format)
	WednesdayCloseTime time.Time `gorm:"column:wednesday_close_time" json:"wednesday_close_time"` // Wednesday closing time (24 hour format)
	// Thursday
	OpenThursday      bool      `gorm:"column:open_thursday" json:"open_thursday"`             // Open on Thursdays (boolean)
	ThursdayOpenTime  time.Time `gorm:"column:thursday_open_time" json:"thursday_open_time"`   // Thursday opening time (24 hour format)
	ThursdayCloseTime time.Time `gorm:"column:thursday_close_time" json:"thursday_close_time"` // Thursday closing time (24 hour format)
	// Friday
	OpenFriday      bool      `gorm:"column:open_friday" json:"open_friday"`             // Open on Fridays (boolean)
	FridayOpenTime  time.Time `gorm:"column:friday_open_time" json:"friday_open_time"`   // Friday opening time (24 hour format)
	FridayCloseTime time.Time `gorm:"column:friday_close_time" json:"friday_close_time"` // Friday closing time (24 hour format)
	// Saturday
	OpenSaturday      bool      `gorm:"column:open_saturday" json:"open_saturday"`             // Open on Saturdays (boolean)
	SaturdayOpenTime  time.Time `gorm:"column:saturday_open_time" json:"saturday_open_time"`   // Saturday opening time (24 hour format)
	SaturdayCloseTime time.Time `gorm:"column:saturday_close_time" json:"saturday_close_time"` // Saturday closing time (24 hour format)
}

// Equal determines if two different Business objects are equal to each other (i.e. all fields match).
//
// Parameters:
// -compareBusiness: The Business object that the calling Business object is being compared to.
//
// Returns:
// -unequalFields []string: The list of fields that did not match between the two Business objects being compared
// -equal bool: If all the fields between the two objects are the same, true is returned. Otherwise, false is returned.
//
// Description:
// This function determines if two Business object instances are equal to each other. The primary purpose of this function
// is to test the functionality of database and handler calls to ensure that the correct objects are being returned and/or
// updated in the database.
func (business *Business) Equal(compareBusiness *Business) (unequalFields []string, equal bool) {
	equal = true

	if business.ID != compareBusiness.ID {
		equal = false
		unequalFields = append(unequalFields, "ID")
	}

	if business.OwnerID != compareBusiness.OwnerID {
		equal = false
		unequalFields = append(unequalFields, "OwnerID")
	}

	if business.MainOfficeID != compareBusiness.MainOfficeID {
		equal = false
		unequalFields = append(unequalFields, "MainOfficeID")
	}

	if business.Name != compareBusiness.Name {
		equal = false
		unequalFields = append(unequalFields, "Name")
	}

	if business.Type != compareBusiness.Type {
		equal = false
		unequalFields = append(unequalFields, "Name")
	}

	if business.CreatedAt.Equal(compareBusiness.CreatedAt) {
		equal = false
		unequalFields = append(unequalFields, "CreatedAt")
	}

	if business.UpdatedAt.Equal(compareBusiness.UpdatedAt) {
		equal = false
		unequalFields = append(unequalFields, "UpdatedAt")
	}

	if !business.DeletedAt.Time.Equal(compareBusiness.DeletedAt.Time) {
		equal = false
		log.Printf("DeletedAt.Time (Business):  %s\nDeletedAt.Time (compareBusiness):  %s", business.DeletedAt.Time, compareBusiness.DeletedAt.Time)
		unequalFields = append(unequalFields, "DeletedAt.Time")
	}

	if business.DeletedAt.Valid != compareBusiness.DeletedAt.Valid {
		equal = false
		log.Printf("DeletedAt.Valid (Business):  %t\nDeletedAt.Valid (compareBusiness):  %t", business.DeletedAt.Valid, compareBusiness.DeletedAt.Valid)
		unequalFields = append(unequalFields, "DeletedAt.Valid")
	}

	return unequalFields, equal
}

// TODO:  Add comment documentation (func CreateBusiness)
func (business *Business) CreateBusiness(db *gorm.DB) (*Business, *Office, error) {
	var office *Office
	// TODO: Add field validation logic (func CreateBusiness) -- add as BeforeCreate gorm hook definition at the top of this file
	if err := db.Create(&business).Error; err != nil {
		return business, office, err
	}

	// Automatically generate a new Office record associated with the business with generic defaults
	var mainOfficeName string = fmt.Sprintf("%s - Main Office", business.Name)
	office.BusinessID = business.ID
	office.ManagerID = business.OwnerID
	office.Name = mainOfficeName

	createdOffice, err := office.CreateOffice(db)
	if err != nil {
		return business, createdOffice, err
	}

	officeIDUpdate := map[string]interface{}{"main_office_id": createdOffice.ID}
	createdBusiness, err := business.UpdateBusiness(db, business.ID, officeIDUpdate)
	if err != nil {
		return createdBusiness, createdOffice, err
	}

	return createdBusiness, createdOffice, nil
}

// TODO:  Add comment documentation (func GetBusiness)
func (business *Business) GetBusiness(db *gorm.DB, businessID uint) (*Business, error) {
	err := db.First(&business, businessID).Error

	if err != nil {
		return business, err
	}

	return business, nil
}

// TODO:  Add comment documentation (func UpdateBusiness)
func (business *Business) UpdateBusiness(db *gorm.DB, businessID uint, updates map[string]interface{}) (*Business, error) {
	// Confirm business exists and get current object
	var err error
	business, err = business.GetBusiness(db, businessID)
	if err != nil {
		return business, err
	}

	// TODO: Add field validation logic (func UpdateBusiness) -- add as BeforeUpdate gorm hook definition at the top of this file

	if err := db.Model(&business).Where("id = ?", businessID).Updates(updates).Error; err != nil {
		return business, err
	}

	return business, nil
}

// TODO:  Add comment documentation (func DeleteBusiness)
func (business *Business) DeleteBusiness(db *gorm.DB, businessID uint) (*Business, error) {
	// Confirm business exists and get current object
	var err error
	business, err = business.GetBusiness(db, businessID)
	if err != nil {
		return business, err
	}

	if err := db.Delete(&business).Error; err != nil {
		return business, err
	}

	return business, nil
}

// TODO:  Add comment documentation (func CreateOffice)
func (office *Office) CreateOffice(db *gorm.DB) (*Office, error) {
	// TODO: Add field validation logic (func CreateOffice) -- add as BeforeCreate gorm hook definition at the top of this file
	if err := db.Create(&office).Error; err != nil {
		return office, err
	}

	return office, nil
}
