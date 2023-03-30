package models

import (
	"log"
	"server/config"
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Service model
// GORM model for all Service records in the database
type Service struct {
	gorm.Model
	OfficeID    uint   `gorm:"column:office_id" json:"office_id"` // ID of Office record Service is associated with
	Name        string `gorm:"column:name" json:"name"`           // Service name
	Description string `gorm:"column:desc" json:"desc"`           // Service description
}

// TODO: Add foreign key logic to ServiceOffering model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
// GORM model for all ServiceOffering records in the database
type ServiceOffering struct {
	gorm.Model
	ServiceID              uint      `gorm:"column:service_id" json:"service_id"`                       // ID of Service record ServiceOffering is associated with
	StaffID                uint      `gorm:"column:staff_id" json:"staff_id"`                           // ID of Staff member that ServiceOffering is associated with
	ResourceID             uint      `gorm:"column:resource_id" json:"resource_id"`                     // ID of Resource that ServiceOffering is associated with
	StartDate              time.Time `gorm:"column:start_date" json:"start_date"`                       // ServiceOffering start date
	EndDate                time.Time `gorm:"column:end_date" json:"end_date"`                           // ServiceOffering end date
	BookingLength          uint      `gorm:"column:booking_length" json:"booking_length"`               // Length of time appointment booking is for (in minutes)
	Price                  uint      `gorm:"column:price" json:"price"`                                 // Price (in cents) for the service being offered
	CancellationFee        uint      `gorm:"column:cancel_fee" json:"cancel_fee"`                       // Fee (in cents) for cancelling appointment after minimum notice cutoff
	MaxConsecutiveBookings uint      `gorm:"column:max_consec_bookings" json:"max_consec_bookings"`     // Max number of consecutive appointments customers can book
	MinCancellationNotice  uint      `gorm:"column:min_cancel_notice" json:"min_cancel_notice"`         // Minimum number of hours appointment cancellation must be made in order to avoid cancellation fee. (null if not applicable, 0 if cancellation fee is always applied)
	MinTimeBetweenClients  uint      `gorm:"column:min_time_betw_clients" json:"min_time_betw_clients"` // Length of time between appointments for differing clients
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
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
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

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file

	err = db.Model(&updateService).Where("id = ?", serviceID).Updates(updates).Error
	returnRecords = map[string]Model{"service": updateService}

	return returnRecords, err
}

// TODO: Cascade delete all records associated with service (ServiceOfferings, etc.)
/*
*Description*

func Delete

Deletes the specified Service record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

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

	// TODO:  Extend delete operations to all of the other object types associated with the Service record as is appropriate (ServiceOfferings, etc.)
	err = db.Delete(deleteService).Error
	returnRecords = map[string]Model{"service": deleteService}

	return returnRecords, err
}

/*
*Description*

func GetID

# Returns ID field from ServiceOffering object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the ServiceOffering object
*/
func (serviceOffering *ServiceOffering) GetID() uint {
	return serviceOffering.ID
}

/*
*Description*

func Create

Creates a new ServiceOffering record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*ServiceOffering>

		The created ServiceOffering record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (serviceOffering *ServiceOffering) Create(db *gorm.DB) (map[string]Model, error) {
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
	err := db.Create(&serviceOffering).Error
	returnRecords := map[string]Model{"service_offering": serviceOffering}
	return returnRecords, err
}

/*
*Description*

func Get

Retrieves an ServiceOffering record in the database by ID if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	serviceOfferingID  <uint>

		The ID of the ServiceOffering record being requested.

*Returns*

	_  <*ServiceOffering>

		The ServiceOffering record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (serviceOffering *ServiceOffering) Get(db *gorm.DB, serviceOfferingID uint) (map[string]Model, error) {
	err := db.First(&serviceOffering, serviceOfferingID).Error
	returnRecords := map[string]Model{"service_offering": serviceOffering}
	return returnRecords, err
}

/*
*Description*

func Update

Updates the specified ServiceOffering record in the database with the specified changes if the record exists.

Returns the updated record along with any errors that are thrown.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve and update the specified record.

	serviceOfferingID  <uint>

		The ID of the ServiceOffering record being updated.

	updates  <map[string]interface{}>

		JSON with the fields that will be updated as keys and the updated values as values.

		Ex:
			{
				"name": "New name",
				"address": "New address"
			}

*Returns*

	_  <*ServiceOffering>

		The ServiceOffering record that is updated in the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (serviceOffering *ServiceOffering) Update(db *gorm.DB, serviceOfferingID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm serviceOfferingID exists in the database and get current object
	returnRecords, err := serviceOffering.Get(db, serviceOfferingID)
	updateServiceOffering := returnRecords["service_offering"]

	if err != nil {
		return returnRecords, err
	}

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file
	err = db.Model(&updateServiceOffering).Where("id = ?", serviceOfferingID).Updates(updates).Error
	returnRecords = map[string]Model{"service_offering": updateServiceOffering}

	return returnRecords, err
}

/*
*Description*

func Delete

Deletes the specified ServiceOffering record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be deleted from.

	serviceOfferingID  <uint>

		The ID of the ServiceOffering record being deleted.

*Returns*

	_  <*ServiceOffering>

		The deleted ServiceOffering record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (serviceOffering *ServiceOffering) Delete(db *gorm.DB, serviceOfferingID uint) (map[string]Model, error) {
	// Confirm serviceOfferingID exists in the database and get current object
	returnRecords, err := serviceOffering.Get(db, serviceOfferingID)
	deleteServiceOffering := returnRecords["service_offering"]

	if err != nil {
		return returnRecords, err
	}

	if config.Debug {
		log.Printf("\n\nServiceOffering object targeted for deletion:\n\n%+v\n\n", deleteServiceOffering)
	}

	err = db.Delete(deleteServiceOffering).Error
	returnRecords = map[string]Model{"service_offering": deleteServiceOffering}

	return returnRecords, err
}
