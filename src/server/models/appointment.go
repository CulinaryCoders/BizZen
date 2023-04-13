package models

import (
	"log"
	"server/config"
	"time"

	"gorm.io/gorm"
)

// TODO: Add foreign key logic to Appointment model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type Appointment struct {
	gorm.Model
	UserID         uint      `gorm:"column:user_id" json:"user_id"`                   // ID of user that booked the appointment
	ServiceID      uint      `gorm:"column:service_id" json:"service_id"`             // ID of service that appointment is for
	Active         bool      `gorm:"column:active" json:"active"`                     // 1 for Active, 0 for Cancelled
	CancelDateTime time.Time `gorm:"column:cancel_date_time" json:"cancel_date_time"` // Date/time when appointment was cancelled (if cancelled, else null)
}

/*
*Description*

func GetID

# Returns ID field from Appointment object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the appointment object
*/
func (appt *Appointment) GetID() uint {
	return appt.ID
}

/*
*Description*

func GetUserID

# Returns UserID field from calling Appointment object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The UserID field from the calling appointment object
*/
func (appt *Appointment) GetUserID() uint {
	return appt.UserID
}

/*
*Description*

func GetServiceID

# Returns ServiceID field from calling Appointment object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ServiceID field from the calling appointment object
*/
func (appt *Appointment) GetServiceID() uint {
	return appt.ServiceID
}

/*
*Description*

func Create

Creates a new Appointment record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*Appointment>

		The created Appointment record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (appt *Appointment) Create(db *gorm.DB) (map[string]Model, error) {
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
	err := db.Create(&appt).Error
	returnRecords := map[string]Model{"appointment": appt}
	return returnRecords, err
}

/*
*Description*

func Get

Retrieves a Appointment record in the database by ID if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	apptID  <uint>

		The ID of the appointment record being requested.

*Returns*

	_  <*Appointment>

		The Appointment record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (appt *Appointment) Get(db *gorm.DB, apptID uint) (map[string]Model, error) {
	err := db.First(&appt, apptID).Error
	returnRecords := map[string]Model{"appointment": appt}
	return returnRecords, err
}

/*
*Description*

func GetAll

Retrieves all Appointment records from the database.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

*Returns*

	_  <[]Appointment>

		The list of Appointment records that are retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (appt *Appointment) GetAll(db *gorm.DB) ([]Appointment, error) {
	var appts []Appointment
	err := db.Find(&appts).Error

	return appts, err
}

// TODO:  Add documentation (func GetRecordListFromSecondaryID)
func (appt *Appointment) GetRecordListFromSecondaryID(db *gorm.DB, secondaryIDJsonKey string, secondaryID uint) ([]Appointment, error) {
	var appts []Appointment

	err := db.Where(map[string]interface{}{secondaryIDJsonKey: secondaryID}).Find(&appts).Error
	return appts, err
}

/*
*Description*

func Update

Updates the specified Appointment record in the database with the specified changes if the record exists.

Returns the updated record along with any errors that are thrown.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve and update the specified record.

	apptID  <uint>

		The ID of the appointment record being updated.

	updates  <map[string]interface{}>

		JSON with the fields that will be updated as keys and the updated values as values.

		Ex:
			{
				"name": "New name",
				"address": "New address"
			}

*Returns*

	_  <*Appointment>

		The Appointment record that is updated in the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (appt *Appointment) Update(db *gorm.DB, apptID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm apptID exists in the database and get current object
	returnRecords, err := appt.Get(db, apptID)
	updateAppointment := returnRecords["appointment"]

	if err != nil {
		return returnRecords, err
	}

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file

	err = db.Model(&updateAppointment).Where("id = ?", apptID).Updates(updates).Error
	returnRecords = map[string]Model{"appointment": updateAppointment}

	return returnRecords, err
}

// TODO: Cascade delete all records associated with appointment (AppointmentOfferings, etc.)
/*
*Description*

func Delete

Deletes the specified Appointment record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

	apptID  <uint>

		The ID of the appointment record being deleted.

*Returns*

	_  <*Appointment>

		The deleted Appointment record.

	_  <error>

		Encountered error (nil if no errors are encountered).

*/
func (appt *Appointment) Delete(db *gorm.DB, apptID uint) (map[string]Model, error) {
	// Confirm apptID exists in the database and get current object
	returnRecords, err := appt.Get(db, apptID)
	deleteAppointment := returnRecords["appointment"]

	if err != nil {
		return returnRecords, err
	}

	if config.Debug {
		log.Printf("\n\nAppointment object targeted for deletion:\n\n%+v\n\n", deleteAppointment)
	}

	// TODO:  Extend delete operations to all of the other object types associated with the Appointment record as is appropriate (AppointmentOfferings, etc.)
	err = db.Delete(deleteAppointment).Error
	returnRecords = map[string]Model{"appointment": deleteAppointment}

	return returnRecords, err
}
