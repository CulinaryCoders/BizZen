package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Business model
type Business struct {
	gorm.Model
	OwnerID      uint   `gorm:"column:owner_id" json:"owner_id"`
	MainOfficeID uint   `gorm:"column:main_office_id" json:"main_office_id"`
	Name         string `gorm:"column:name" json:"name"`
	Type         string `gorm:"column:type" json:"type"`
}

// TODO: Add foreign key logic to Office model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type Office struct {
	gorm.Model
	BusinessID    uint      `gorm:"column:business_id" json:"business_id"`
	ContactInfoID uint      `gorm:"column:contact_info_id" json:"contact_info_id"`
	ManagerID     uint      `gorm:"column:manager_id" json:"manager_id"`
	Name          string    `gorm:"not null;column:name" json:"name"`
	OpeningTime   time.Time `gorm:"column:open_time" json:"open_time"`
	ClosingTime   time.Time `gorm:"column:close_time" json:"close_time"`
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
func (business *Business) CreateBusiness(db *gorm.DB) (*Business, error) {
	// TODO: Add field validation logic (func CreateBusiness) -- add as BeforeCreate gorm hook definition at the top of this file
	if err := db.Create(&business).Error; err != nil {
		return business, err
	}

	return business, nil
}

// TODO:  Add comment documentation (func GetBusiness)
func (business *Business) GetBusiness(db *gorm.DB, businessID string) (*Business, error) {
	err := db.First(&business, businessID).Error

	if err != nil {
		return business, err
	}

	return business, nil
}

// TODO:  Add comment documentation (func UpdateBusiness)
func (business *Business) UpdateBusiness(db *gorm.DB, businessID string, updates map[string]interface{}) (*Business, error) {
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
func (business *Business) DeleteBusiness(db *gorm.DB, businessID string) (*Business, error) {
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
