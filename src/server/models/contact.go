package models

import (
	"log"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to ContactInfo model
type ContactInfo struct {
	gorm.Model
	OwnerID      uint   `gorm:"column:owner_id" json:"owner_id"`
	AddressID    uint   `gorm:"column:address_id" json:"address_id"`
	PhoneNumber1 string `gorm:"column:phone1" json:"phone1"`
	PhoneNumber2 string `gorm:"column:phone2" json:"phone2"`
	FaxNumber    string `gorm:"column:fax" json:"fax"`
}

type Address struct {
	gorm.Model
	Address1 string `gorm:"not null;column:address1" json:"address1"`
	Address2 string `gorm:"column:address2" json:"address2"`
	City     string `gorm:"not null;column:city" json:"city"`
	State    string `gorm:"not null;column:state" json:"state"`
	ZipCode  string `gorm:"not null;column:zip" json:"zip"`
}

func (address *Address) Equal(compareAddress *Address) (unequalFields []string, equal bool) {
	equal = true

	if address.ID != compareAddress.ID {
		equal = false
		unequalFields = append(unequalFields, "ID")
	}

	if address.Address1 != compareAddress.Address1 {
		equal = false
		unequalFields = append(unequalFields, "Address1")
	}

	if address.Address2 != compareAddress.Address2 {
		equal = false
		unequalFields = append(unequalFields, "Address2")
	}

	if address.City != compareAddress.City {
		equal = false
		unequalFields = append(unequalFields, "City")
	}

	if address.State != compareAddress.State {
		equal = false
		unequalFields = append(unequalFields, "State")
	}

	if address.ZipCode != compareAddress.ZipCode {
		equal = false
		unequalFields = append(unequalFields, "ZipCode")
	}

	if address.CreatedAt.Equal(compareAddress.CreatedAt) {
		equal = false
		unequalFields = append(unequalFields, "CreatedAt")
	}

	if address.UpdatedAt.Equal(compareAddress.UpdatedAt) {
		equal = false
		unequalFields = append(unequalFields, "UpdatedAt")
	}

	if !address.DeletedAt.Time.Equal(compareAddress.DeletedAt.Time) {
		equal = false
		log.Printf("DeletedAt.Time (Address):  %s\nDeletedAt.Time (compareAddress):  %s", address.DeletedAt.Time, compareAddress.DeletedAt.Time)
		unequalFields = append(unequalFields, "DeletedAt.Time")
	}

	if address.DeletedAt.Valid != compareAddress.DeletedAt.Valid {
		equal = false
		log.Printf("DeletedAt.Valid (Address):  %t\nDeletedAt.Valid (compareAddress):  %t", address.DeletedAt.Valid, compareAddress.DeletedAt.Valid)
		unequalFields = append(unequalFields, "DeletedAt.Valid")
	}

	return unequalFields, equal
}

// TODO:  Add comment documentation (func CreateAddress)
func (address *Address) CreateAddress(db *gorm.DB) (*Address, error) {
	// TODO: Add field validation logic (func CreateAddress) -- add as BeforeCreate gorm hook definition at the top of this file
	if err := db.Create(&address).Error; err != nil {
		return address, err
	}

	return address, nil
}

// TODO:  Add comment documentation (func GetAddress)
func (address *Address) GetAddress(db *gorm.DB, addressID uint) (*Address, error) {
	err := db.First(&address, addressID).Error

	if err != nil {
		return address, err
	}

	return address, nil
}

// TODO:  Add comment documentation (func UpdateAddress)
func (address *Address) UpdateAddress(db *gorm.DB, addressID uint, updates map[string]interface{}) (*Address, error) {
	// Confirm address exists and get current object
	var err error
	address, err = address.GetAddress(db, addressID)
	if err != nil {
		return address, err
	}

	// TODO: Add field validation logic (func UpdateAddress) -- add as BeforeUpdate gorm hook definition at the top of this file

	if err := db.Model(&address).Where("id = ?", addressID).Updates(updates).Error; err != nil {
		return address, err
	}

	return address, nil
}

// TODO:  Add comment documentation (func DeleteAddress)
func (address *Address) DeleteAddress(db *gorm.DB, addressID uint) (*Address, error) {
	// Confirm address exists and get current object
	var err error
	address, err = address.GetAddress(db, addressID)
	if err != nil {
		return address, err
	}

	if err := db.Delete(&address).Error; err != nil {
		return address, err
	}

	return address, nil
}
