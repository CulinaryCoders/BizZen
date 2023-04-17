package models

import (
	"log"
	"server/config"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TODO: Add foreign key logic to Appointment model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
type Appointment struct {
	gorm.Model
	UserID         uint       `gorm:"column:user_id" json:"user_id"`                                // ID of user that booked the appointment
	ServiceID      uint       `gorm:"column:service_id" json:"service_id"`                          // ID of service that appointment is for
	Active         bool       `gorm:"column:active;default:true" json:"active"`                     // 1 for Active, 0 for Cancelled
	CancelDateTime *time.Time `gorm:"column:cancel_date_time;default:null" json:"cancel_date_time"` // Date/time when appointment was cancelled (if cancelled, else null)
}

// TODO:  Add documentation for GORM db hook (func AfterCreate)
func (appt *Appointment) AfterCreate(db *gorm.DB) error {
	if config.Debug {
		log.Println("AfterCreate hook executed [Appointment model].")
	}

	var serviceIDJsonKey string = "service_id"
	appts, err := appt.GetRecordsBySecondaryID(db, serviceIDJsonKey, appt.ServiceID)
	if err != nil {
		return err
	}

	service := Service{}
	err = db.Model(Service{}).Where(appt.ServiceID).Find(&service).Error
	if err != nil {
		return err
	}

	if config.Debug {
		log.Printf("Original Service:\n\n%v\n\n", service)
	}

	var active_appt_ct int = 0
	for _, appt := range appts {
		if appt.Active {
			active_appt_ct++
		}
	}
	updates := map[string]interface{}{
		"appt_ct": active_appt_ct,
		"is_full": active_appt_ct == int(service.Capacity),
	}

	updatedService, err := service.Update(db, appt.ServiceID, updates)
	if err != nil {
		return err
	}

	if config.Debug {
		log.Printf("Updated Service:\n\n%v\n\n", updatedService["service"])
	}

	return nil
}

// TODO:  Add documentation for GORM db hook (func AfterUpdate)
func (appt *Appointment) AfterUpdate(db *gorm.DB) error {
	if config.Debug {
		log.Println("AfterCreate hook executed [Appointment model].")
	}

	var serviceIDJsonKey string = "service_id"
	appts, err := appt.GetRecordsBySecondaryID(db, serviceIDJsonKey, appt.ServiceID)
	if err != nil {
		return err
	}

	service := Service{}
	err = db.Model(Service{}).Where(appt.ServiceID).Find(&service).Error
	if err != nil {
		return err
	}

	if config.Debug {
		log.Printf("Appointment:\n\n%v\n\n", appt)
		log.Printf("Original Service:\n\n%v\n\n", service)
	}

	var active_appt_ct int = 0
	for _, appt := range appts {
		if appt.Active {
			active_appt_ct++
		}
	}
	updates := map[string]interface{}{
		"appt_ct": active_appt_ct,
		"is_full": active_appt_ct == int(service.Capacity),
	}

	updatedService, err := service.Update(db, appt.ServiceID, updates)
	if err != nil {
		return err
	}

	if config.Debug {
		log.Printf("Updated Service:\n\n%v\n\n", updatedService["service"])
	}

	return nil
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

/*
*Description*

func GetRecordsBySecondaryID

Retrieves a list of Appointment records from the database that are associated with the specified secondary key.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

	secondaryIDJsonKey  <string>

		The JSON key for the secondary ID attribute.

	secondaryID  <uint>

		The secondary ID value.

*Returns*

	_  <[]Appointment>

		The list of Appointment records that are retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (appt *Appointment) GetRecordsBySecondaryID(db *gorm.DB, secondaryIDJsonKey string, secondaryID uint) ([]Appointment, error) {
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

	err = db.Model(&updateAppointment).Clauses(clause.Returning{}).Where("id = ?", apptID).Updates(updates).Error
	returnRecords = map[string]Model{"appointment": updateAppointment}

	return returnRecords, err
}

// TODO:  Add documentation (func Cancel)
func (appt *Appointment) Cancel(db *gorm.DB, apptID uint) (map[string]Model, error) {
	var updates map[string]interface{} = map[string]interface{}{
		"active":           false,
		"cancel_date_time": time.Now(),
	}
	returnedRecords, err := appt.Update(db, apptID, updates)
	log.Println(time.Now())
	return returnedRecords, err
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
