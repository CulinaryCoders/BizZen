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

/*
*Description*

func Equal

Determines if two different Business objects are equal to each other (i.e. all fields match).

The primary purpose of this function is to test the functionality of database and handler calls to ensure that
the correct objects are being returned and/or updated in the database.

*Parameters*

	compareBusiness  <*Business>

		The Business object that the calling Business object is being compared to

*Returns*

	unequalFields  <[]string>

		The list of fields that did not match between the two Business objects being compared

	equal  <bool>

		If all the fields between the two objects are the same, true is returned. Otherwise, false is returned.

*Response format*

	N/A (None)
*/
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

/*
*Description*

func getID

# Returns ID field from Business object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the business object
*/
func (business *Business) getID() uint {
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

*Response format*

	N/A (None)
*/
func (business *Business) Create(db *gorm.DB) (map[string]Model, error) {
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
	err := db.Create(&business).Error
	if err != nil {
		returnRecords := map[string]Model{"business": business, "office": &Office{}}
		return returnRecords, err
	}

	// Automatically generate a new Office record associated with the business with generic defaults
	var mainOfficeName string = fmt.Sprintf("%s - Main Office", business.Name)

	office := Office{
		BusinessID: business.ID,
		ManagerID:  business.OwnerID,
		Name:       mainOfficeName,
	}

	returnRecords, err := office.Create(db)
	createdOffice := returnRecords["office"]

	returnRecords = map[string]Model{"business": business, "office": createdOffice}

	if err != nil {
		log.Panicln("Could not create new Office record.")

		return returnRecords, err
	}

	//  Set MainOfficeID for new Business record to new Office record that was created
	officeIDUpdate := map[string]interface{}{"main_office_id": createdOffice.getID()}
	returnRecords, err = business.Update(db, business.ID, officeIDUpdate)
	updatedBusiness := returnRecords["business"]

	returnRecords = map[string]Model{"business": updatedBusiness, "office": createdOffice}

	if err != nil {
		log.Panicln("Could not update Office ID for Business object.")
	}

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

*Response format*

	N/A (None)
*/
func (business *Business) Get(db *gorm.DB, businessID uint) (map[string]Model, error) {
	err := db.First(&business, businessID).Error
	returnRecords := map[string]Model{"business": business}
	return returnRecords, err
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

*Response format*

	N/A (None)
*/
func (business *Business) Update(db *gorm.DB, businessID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm businessID exists in the database and get current object
	returnRecords, err := business.Get(db, businessID)
	updateBusiness := returnRecords["business"]

	if err != nil {
		return returnRecords, err
	}

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file

	err = db.Model(&updateBusiness).Where("id = ?", businessID).Updates(updates).Error
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

*Response format*

	N/A (None)
*/
func (business *Business) Delete(db *gorm.DB, businessID uint) (map[string]Model, error) {
	// Confirm businessID exists in the database and get current object
	returnRecords, err := business.Get(db, businessID)
	deleteBusiness := returnRecords["business"]

	if err != nil {
		return returnRecords, err
	}

	// TODO:  Extend delete operations to all of the other object types associated with the Business record as is appropriate (Offices, Services, etc.)
	err = db.Delete(&deleteBusiness).Error
	returnRecords = map[string]Model{"business": deleteBusiness}

	return returnRecords, err
}

/*
*Description*

func getID

# Returns ID field from Business object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the business object
*/
func (office *Office) getID() uint {
	return office.ID
}

/*
*Description*

func Create

Creates a new Office record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*Office>

		The created Office record.

	_  <error>

		Encountered error (nil if no errors are encountered).

*Response format*

	N/A (None)
*/
func (office *Office) Create(db *gorm.DB) (map[string]Model, error) {
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
	err := db.Create(&office).Error
	returnRecords := map[string]Model{"office": office}
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

*Response format*

	N/A (None)
*/
func (office *Office) Get(db *gorm.DB, officeID uint) (map[string]Model, error) {
	err := db.First(&office, officeID).Error
	returnRecords := map[string]Model{"office": office}
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

*Response format*

	N/A (None)
*/
func (office *Office) Update(db *gorm.DB, officeID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm officeID exists in the database and get current object
	returnRecords, err := office.Get(db, officeID)
	updateOffice := returnRecords["office"]

	if err != nil {
		return returnRecords, err
	}

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file
	err = db.Model(&updateOffice).Where("id = ?", officeID).Updates(updates).Error
	returnRecords = map[string]Model{"office": updateOffice}

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

*Response format*

	N/A (None)
*/
func (office *Office) Delete(db *gorm.DB, officeID uint) (map[string]Model, error) {
	// Confirm address exists and get current object
	returnRecords, err := office.Get(db, officeID)
	deleteOffice := returnRecords["office"]

	if err != nil {
		return returnRecords, err
	}

	err = db.Delete(&office).Error
	returnRecords = map[string]Model{"office": deleteOffice}

	return returnRecords, err
}
