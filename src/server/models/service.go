package models

import (
	"errors"
	"fmt"
	"log"
	"server/config"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM model for all Service records in the database
type Service struct {
	gorm.Model
	BusinessID    uint      `gorm:"column:business_id" json:"business_id"`         // ID of Business that Service is associated with
	Name          string    `gorm:"column:name" json:"name"`                       // Service name
	Description   string    `gorm:"column:desc" json:"desc"`                       // Service description
	StartDateTime time.Time `gorm:"column:start_date_time" json:"start_date_time"` // Date/time that the service starts
	Length        uint      `gorm:"column:length" json:"length"`                   // Length of time in minutes that the service will take
	Capacity      uint      `gorm:"column:capacity" json:"capacity"`               // Number of users that can sign up for the service
	CancelFee     uint      `gorm:"column:cancel_fee" json:"cancel_fee"`           // Fee (in cents) for cancelling appointment after minimum notice cutoff
	Price         uint      `gorm:"column:price" json:"price"`                     // Price (in cents) for the service being offered
	AppointmentCt int       `gorm:"column:appt_ct" json:"appt_ct" default:"0"`     // Number of active appointments scheduled for the Service
	IsFull        bool      `gorm:"column:is_full" json:"is_full" default:"false"` // True if number of active appointments equals the capacity for the Service (False if not)
}

/*
*Description*

func AfterDelete (GORM hook)

Deletes all of the Appointment records in the database that are associated with a Service record
when the Service record is deleted.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (service *Service) AfterDelete(db *gorm.DB) error {
	appt := Appointment{}
	var serviceIDJsonKey string = "service_id"

	_, err := appt.DeleteRecordsBySecondaryID(db, serviceIDJsonKey, service.ID)
	if err != nil {
		return err
	}

	return nil
}

/*
*Description*

func IDExists

Checks to see if a Service record with the specified ID already exists in the database.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be queried for the specified Service ID.

	serviceID  <uint>

		The Service ID to check for.

*Returns*

	_  <bool>

		'true' if a Service record exists in the database with the specified ID. 'false' if not.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (service *Service) IDExists(db *gorm.DB, serviceID uint) (bool, error) {
	var idExists bool
	err := db.Model(Service{}).Select("count(*) > 0").Where("id = ?", serviceID).Find(&idExists).Error
	return idExists, err
}

/*
*Description*

func GetID

# Returns ID field from Service object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the service object
*/
func (service *Service) GetID() uint {
	return service.ID
}

/*
*Description*

func Create

Creates a new Service record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*Service>

		The created Service record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (service *Service) Create(db *gorm.DB) (map[string]Model, error) {
	err := db.Create(&service).Error
	returnRecords := map[string]Model{"service": service}
	return returnRecords, err
}

/*
*Description*

func Get

Retrieves a Service record in the database by ID if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	serviceID  <uint>

		The ID of the service record being requested.

*Returns*

	_  <*Service>

		The Service record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) Get(db *gorm.DB, serviceID uint) (map[string]Model, error) {
	err := db.First(&service, serviceID).Error
	returnRecords := map[string]Model{"service": service}
	return returnRecords, err
}

/*
*Description*

func GetAll

Retrieves all Service records from the database.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

*Returns*

	_  <[]Service>

		The list of Service records that are retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) GetAll(db *gorm.DB) ([]Service, error) {
	var services []Service
	err := db.Find(&services).Error

	return services, err
}

/*
*Description*

func GetRecordsByPrimaryIDs

Retrieves a list of Service records from the database using their IDs (primary key).

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

	ids  <[]uint>

		The list of Service IDs that will be used to retrieve Service records.

*Returns*

	_  <[]Service>

		The list of Service records that are retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) GetRecordsByPrimaryIDs(db *gorm.DB, ids []uint) ([]Service, error) {
	var services []Service

	err := db.Where(ids).Find(&services).Error
	return services, err
}

/*
*Description*

func GetRecordsBySecondaryID

Retrieves a list of Service records from the database that are associated with the specified secondary key.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

	secondaryIDJsonKey  <string>

		The JSON key for the secondary ID attribute.

	secondaryID  <uint>

		The secondary ID value.

*Returns*

	_  <[]Service>

		The list of Service records that are retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) GetRecordsBySecondaryID(db *gorm.DB, secondaryIDJsonKey string, secondaryID uint) ([]Service, error) {
	var services []Service

	err := db.Where(map[string]interface{}{secondaryIDJsonKey: secondaryID}).Find(&services).Error
	return services, err
}

/*
*Description*

func GetAppointments

Retrieves the list of all Appointments that are associated with the specified Service.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

	serviceID  <uint>

		The Service ID that will be used to retrieve the list of Appointment records.

*Returns*

	_  <[]Appointment>

		The list of Appointment records that are retrieved from the database that are associated with the specified Service.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) GetAppointments(db *gorm.DB, serviceID uint) ([]Appointment, error) {
	var appt Appointment
	var appts []Appointment

	// Confirm Service record exists for specified ID
	_, err := service.Get(db, serviceID)
	if err != nil {
		var errorMessage string = fmt.Sprintf("Service ID (%d) does not exist in the database.  [%s]", serviceID, err)
		return appts, errors.New(errorMessage)
	}

	// Get list of appointments for specified ServiceID
	var serviceIDJsonKey string = "service_id"
	appts, err = appt.GetRecordsBySecondaryID(db, serviceIDJsonKey, serviceID)

	return appts, err
}

/*
*Description*

func GetUsers

Retrieves the list of all the Users that have signed up for a particular Service.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

	serviceID  <uint>

		The Service ID that will be used to retrieve the list of User records.

*Returns*

	_  <[]User>

		The list of User records that are retrieved from the database that have an appointment scheduled for the specified Service.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) GetUsers(db *gorm.DB, serviceID uint, activeOnly bool) ([]User, error) {
	var apptsUserIDs []uint
	var user User
	var users []User

	// Get list of appointments for specified ServiceID
	appts, err := service.GetAppointments(db, serviceID)
	if err != nil {
		return users, err
	}

	// Get list of UserIDs from appointments
	for _, appt := range appts {
		if !activeOnly || (activeOnly && appt.Active) {
			apptsUserIDs = append(apptsUserIDs, appt.GetUserID())
		}
	}

	// Get list of Users from appointment UserIDs
	if len(apptsUserIDs) > 0 {
		users, err = user.GetRecordsByPrimaryIDs(db, apptsUserIDs)
		if err != nil {
			return users, err
		}
	}

	return users, nil
}

/*
*Description*

func Update

Updates the specified Service record in the database with the specified changes if the record exists.

Returns the updated record along with any errors that are thrown.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve and update the specified record.

	serviceID  <uint>

		The ID of the service record being updated.

	updates  <map[string]interface{}>

		JSON with the fields that will be updated as keys and the updated values as values.

		Ex:
			{
				"name": "New name",
				"address": "New address"
			}

*Returns*

	_  <*Service>

		The Service record that is updated in the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) Update(db *gorm.DB, serviceID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm serviceID exists in the database and get current object
	returnRecords, err := service.Get(db, serviceID)
	updateService := returnRecords["service"]

	if err != nil {
		return returnRecords, err
	}

	err = db.Model(&updateService).Clauses(clause.Returning{}).Where("id = ?", serviceID).Updates(updates).Error
	returnRecords = map[string]Model{"service": updateService}

	return returnRecords, err
}

/*
*Description*

func Delete

Deletes the specified Service record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be deleted.

	serviceID  <uint>

		The ID of the service record being deleted.

*Returns*

	_  <*Service>

		The deleted Service record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (service *Service) Delete(db *gorm.DB, serviceID uint) (map[string]Model, error) {
	// Confirm serviceID exists in the database and get current object
	returnRecords, err := service.Get(db, serviceID)
	deleteService := returnRecords["service"]

	if err != nil {
		return returnRecords, err
	}

	if config.Debug {
		log.Printf("\n\nService object targeted for deletion:\n\n%+v\n\n", deleteService)
	}

	err = db.Delete(deleteService).Error
	returnRecords = map[string]Model{"service": deleteService}

	return returnRecords, err
}

/*
*Description*

func DeleteRecordsBySecondaryID

Deletes a list of Service records from the database that are associated with the specified secondary key.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

	secondaryIDJsonKey  <string>

		The JSON key for the secondary ID attribute.

	secondaryID  <uint>

		The secondary ID value.

*Returns*

	_  <[]Appointment>

		The list of Service records that were deleted from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (service *Service) DeleteRecordsBySecondaryID(db *gorm.DB, secondaryIDJsonKey string, secondaryID uint) ([]Model, error) {
	var services []Service
	var deletedServices []Model

	services, err := service.GetRecordsBySecondaryID(db, secondaryIDJsonKey, secondaryID)
	if err != nil {
		return deletedServices, err
	}

	for _, deleteService := range services {
		returnedRecords, err := deleteService.Delete(db, deleteService.ID)
		if err != nil {
			return deletedServices, err
		}

		deletedServices = append(deletedServices, returnedRecords["service"])
	}

	return deletedServices, nil
}
