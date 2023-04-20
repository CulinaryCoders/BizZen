package models

import (
	"errors"
	"fmt"
	"log"

	"server/config"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM model for all Business records in the database
type Business struct {
	gorm.Model
	OwnerID uint   `gorm:"column:owner_id" json:"owner_id"` // ID of User account that owns the business record
	Name    string `gorm:"column:name" json:"name"`         // Business name
}

/*
*Description*

func AfterDelete (GORM hook)

Deletes all of the Service records in the database that are associated with a Business record
when the Business record is deleted.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the operations will be performed.

*Returns*

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (business *Business) AfterDelete(db *gorm.DB) error {
	service := Service{}
	var businessIDJsonKey string = "business_id"

	_, err := service.DeleteRecordsBySecondaryID(db, businessIDJsonKey, business.ID)
	if err != nil {
		return err
	}

	return nil
}

/*
*Description*

func IDExists

Checks to see if a Business record with the specified ID already exists in the database.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be queried for the specified Business ID.

	businessID  <uint>

		The Business ID to check for.

*Returns*

	_  <bool>

		'true' if a Business record exists in the database with the specified ID. 'false' if not.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (business *Business) IDExists(db *gorm.DB, businessID uint) (bool, error) {
	var idExists bool
	err := db.Model(Business{}).Select("count(*) > 0").Where("id = ?", businessID).Find(&idExists).Error
	return idExists, err
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

func GetServiceAppointments

Retrieves the list of all Appointments (and the Service each Appointment is for) that are associated with the specified User.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

	userID  <uint>

		The User ID that will be used to retrieve the list of Appointment/Service records.

*Returns*

	_  <[]map[string]interface{}>

		A list of JSON objects that each have an "appointment" key and a "service" key with the respective Appointment/Service record.

		Ex:
			[
				{
					"service": {
						"ID": 22,
						"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
						"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
						"DeletedAt": null,
						"business_id":42,
						"name":"Yoga class",
						"desc":"30 minute beginner yoga class",
						"start_date_time":"2023-05-31T14:30:00.0000000-05:00",
						"length":30,
						"capacity":20,
						"price":2000,
						"cancel_fee":0,
						"appt_ct":13,
						"is_full":false
					},
					"appointments": [
						{
							"ID":11,
							"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
							"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
							"DeletedAt": null,
							"service_id":22,
							"user_id":33,
							"cancel_date_time":null,
							"active":true
						},
						{
							"ID":51,
							"CreatedAt": "2020-05-21T01:23:45.6789012-05:00",
							"UpdatedAt": "2020-05-21T01:23:45.6789012-05:00",
							"DeletedAt": null,
							"service_id":22,
							"user_id":26,
							"cancel_date_time":null,
							"active":true
						},
						...
					]
				},
				{
					"service": {
						"ID": 55,
						"CreatedAt": "2020-02-05T01:23:45.6789012-05:00",
						"UpdatedAt": "2020-02-05T01:23:45.6789012-05:00",
						"DeletedAt": null,
						"business_id":99,
						"name":"Spin class",
						"desc":"60 minute intermediate spin class",
						"start_date_time":"2023-04-20T10:00:00.0000000-05:00",
						"length":60,
						"capacity":10,
						"price":5000,
						"cancel_fee":1000,
						"appt_ct":5,
						"is_full":false
					},
					"appointments": [
						{
							"ID":44,
							"CreatedAt": "2020-01-01T01:23:45.6789012-05:00",
							"UpdatedAt": "2020-01-01T01:23:45.6789012-05:00",
							"DeletedAt": null,
							"service_id":55,
							"user_id":66,
							"cancel_date_time":null,
							"active":true
						},
						{
							"ID":69,
							"CreatedAt": "2020-03-14T16:47:51.1387974-05:00",
							"UpdatedAt": "2020-03-14T16:47:51.1387974-05:00",
							"DeletedAt": null,
							"service_id":55,
							"user_id":85,
							"cancel_date_time":null,
							"active":true
						},
						...
					]
				},
				...
			]

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (business *Business) GetServiceAppointments(db *gorm.DB, businessID uint, activeOnly bool) ([]map[string]interface{}, error) {
	service := Service{}
	var services []Service
	var serviceAppointments []map[string]interface{}

	// Get list of appointments for specified UserID
	var businessIDJsonKey string = "business_id"
	services, svcErr := service.GetRecordsBySecondaryID(db, businessIDJsonKey, businessID)
	if svcErr != nil {
		var errorMessage string = fmt.Sprintf("Business ID (%d) does not have any service records in the database.  [%s]", businessID, svcErr)
		return serviceAppointments, errors.New(errorMessage)
	}

	// Get list of ServiceIDs from user's appointments
	var serviceIDJsonKey string = "service_id"
	for _, service := range services {
		// Get Appointments associated with each of the business' services
		serviceAppt := Appointment{}
		serviceID := service.GetID()
		appts, apptErr := serviceAppt.GetRecordsBySecondaryID(db, serviceIDJsonKey, serviceID)
		if apptErr != nil {
			return serviceAppointments, apptErr
		}

		var finalApptList []Appointment
		if activeOnly {
			for _, appt := range appts {
				if appt.Active {
					finalApptList = append(finalApptList, appt)
				}
			}
		} else {
			finalApptList = appts
		}

		// Structure JSON appropriately and append to list of service appointments
		var svcAppt map[string]interface{} = map[string]interface{}{"service": service, "appointments": finalApptList}
		serviceAppointments = append(serviceAppointments, svcAppt)
	}

	return serviceAppointments, nil
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

	err = db.Model(&updateBusiness).Clauses(clause.Returning{}).Where("id = ?", businessID).Updates(updates).Error
	returnRecords = map[string]Model{"business": updateBusiness}

	return returnRecords, err
}

/*
*Description*

func Delete

Deletes the specified Business record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be deleted.

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

	err = db.Delete(deleteBusiness).Error
	returnRecords = map[string]Model{"business": deleteBusiness}

	return returnRecords, err
}
