package models

import (
	"log"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to ContactInfo model
// GORM model for all ContactInfo records in the database
type ContactInfo struct {
	gorm.Model
	OwnerID      uint   `gorm:"column:owner_id" json:"owner_id"`     // ID of User record that the ContactInfo record is associated with
	AddressID    uint   `gorm:"column:address_id" json:"address_id"` // ID of Address record associated with the ContactInfo record
	PhoneNumber1 string `gorm:"column:phone1" json:"phone1"`         // Primary phone number
	PhoneNumber2 string `gorm:"column:phone2" json:"phone2"`         // Secondary phone number
	FaxNumber    string `gorm:"column:fax" json:"fax"`               // Fax number
}

// GORM model for all Address records in the database
type Address struct {
	gorm.Model
	Address1 string `gorm:"not null;column:address1" json:"address1"` // Address line 1
	Address2 string `gorm:"column:address2" json:"address2"`          // Address line 2
	City     string `gorm:"not null;column:city" json:"city"`         // City
	State    string `gorm:"not null;column:state" json:"state"`       // State (2 letter abbreviation)
	ZipCode  string `gorm:"not null;column:zip" json:"zip"`           // Zip code
}

/*
*Description*

func Equal

Determines if two different Address objects are equal to each other (i.e. all fields match).

The primary purpose of this function is to test the functionality of database and handler calls to ensure that
the correct objects are being returned and/or updated in the database.

*Parameters*

	compareAddress  <*Address>

		The Address object that the calling Address object is being compared to

*Returns*

	unequalFields  <[]string>

		The list of fields that did not match between the two Address objects being compared

	equal  <bool>

		If all the fields between the two objects are the same, true is returned. Otherwise, false is returned.
*/
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

/*
*Description*

func getID

# Returns ID field from Address object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the address object
*/
func (address *Address) getID() uint {
	return address.ID
}

/*
*Description*

func Create

Creates a new Address record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*Address>

		The created Address record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (address *Address) Create(db *gorm.DB) (map[string]Model, error) {
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
	err := db.Create(&address).Error
	returnRecords := map[string]Model{"address": address}
	return returnRecords, err
}

/*
*Description*

func Get

Retrieves an Address record in the database by ID if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	addressID  <uint>

		The ID of the address record being requested.

*Returns*

	_  <*Address>

		The Address record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (address *Address) Get(db *gorm.DB, addressID uint) (map[string]Model, error) {
	err := db.First(&address, addressID).Error
	returnRecords := map[string]Model{"address": address}
	return returnRecords, err
}

/*
*Description*

func Update

Updates the specified Address record in the database with the specified changes if the record exists.

Returns the updated record along with any errors that are thrown.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve and update the specified record.

	addressID  <uint>

		The ID of the address record being updated.

	updates  <map[string]interface{}>

		JSON with the fields that will be updated as keys and the updated values as values.

		Ex:
			{
				"name": "New name",
				"address": "New address"
			}

*Returns*

	_  <*Address>

		The Address record that is updated in the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (address *Address) Update(db *gorm.DB, addressID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm addressID exists in the database and get current object
	returnRecords, err := address.Get(db, addressID)
	updateAddress := returnRecords["address"]

	if err != nil {
		return returnRecords, err
	}

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file
	err = db.Model(&updateAddress).Where("id = ?", addressID).Updates(updates).Error
	returnRecords = map[string]Model{"address": updateAddress}

	return returnRecords, err
}

/*
*Description*

func Delete

Deletes the specified Address record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be deleted from.

	addressID  <uint>

		The ID of the address record being deleted.

*Returns*

	_  <*Address>

		The deleted Address record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (address *Address) Delete(db *gorm.DB, addressID uint) (map[string]Model, error) {
	// Confirm addressID exists in the database and get current object
	returnRecords, err := address.Get(db, addressID)
	deleteAddress := returnRecords["address"]

	if err != nil {
		return returnRecords, err
	}

	err = db.Delete(&address).Error
	returnRecords = map[string]Model{"address": deleteAddress}

	return returnRecords, err
}
