package models

import (
	"log"

	"server/config"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TODO: Add foreign key logic to Business model
// GORM model for all Business records in the database
type Business struct {
	gorm.Model
	OwnerID uint   `gorm:"column:owner_id" json:"owner_id"` // ID of User account that owns the business record
	Name    string `gorm:"column:name" json:"name"`         // Business name
}

/*
*Description*

func GetID

# Returns ID field from Business object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the business object
*/
func (business *Business) GetID() uint {
	return business.ID
}

/*
*Description*

func Create

Creates a new Business record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*Business>

		The created Business record.

	_  <*Office>

		The created Office record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (business *Business) Create(db *gorm.DB) (map[string]Model, error) {
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
	err := db.Create(&business).Error
	if err != nil {
		returnRecords := map[string]Model{"business": business}
		return returnRecords, err
	}

	returnRecords := map[string]Model{"business": business}

	return returnRecords, err
}

/*
*Description*

func Get

Retrieves a Business record in the database by ID if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	businessID  <uint>

		The ID of the business record being requested.

*Returns*

	_  <*Business>

		The Business record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (business *Business) Get(db *gorm.DB, businessID uint) (map[string]Model, error) {
	err := db.First(&business, businessID).Error
	returnRecords := map[string]Model{"business": business}
	return returnRecords, err
}

/*
*Description*

func GetAll

Retrieves all Business records from the database.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

*Returns*

	_  <[]Business>

		The list of Business records that are retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (business *Business) GetAll(db *gorm.DB) ([]Business, error) {
	var businesses []Business
	err := db.Find(&businesses).Error

	return businesses, err
}

/*
*Description*

func Update

Updates the specified Business record in the database with the specified changes if the record exists.

Returns the updated record along with any errors that are thrown.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve and update the specified record.

	businessID  <uint>

		The ID of the business record being updated.

	updates  <map[string]interface{}>

		JSON with the fields that will be updated as keys and the updated values as values.

		Ex:
			{
				"name": "New name",
				"address": "New address"
			}

*Returns*

	_  <*Business>

		The Business record that is updated in the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (business *Business) Update(db *gorm.DB, businessID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm businessID exists in the database and get current object
	returnRecords, err := business.Get(db, businessID)
	updateBusiness := returnRecords["business"]

	if err != nil {
		return returnRecords, err
	}

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file

	err = db.Model(&updateBusiness).Clauses(clause.Returning{}).Where("id = ?", businessID).Updates(updates).Error
	returnRecords = map[string]Model{"business": updateBusiness}

	return returnRecords, err
}

// TODO: Cascade delete all records associated with business (operating hours, offices, contact info, etc.)
/*
*Description*

func Delete

Deletes the specified Business record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

	businessID  <uint>

		The ID of the business record being deleted.

*Returns*

	_  <*Business>

		The deleted Business record.

	_  <error>

		Encountered error (nil if no errors are encountered).

*/
func (business *Business) Delete(db *gorm.DB, businessID uint) (map[string]Model, error) {
	// Confirm businessID exists in the database and get current object
	returnRecords, err := business.Get(db, businessID)
	deleteBusiness := returnRecords["business"]

	if err != nil {
		return returnRecords, err
	}

	if config.Debug {
		log.Printf("\n\nBusiness object targeted for deletion:\n\n%+v\n\n", deleteBusiness)
	}

	// TODO:  Extend delete operations to all of the other object types associated with the Business record as is appropriate (Offices, Services, etc.)
	err = db.Delete(deleteBusiness).Error
	returnRecords = map[string]Model{"business": deleteBusiness}

	return returnRecords, err
}
