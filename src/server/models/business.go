package models

import (
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
