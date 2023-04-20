package models

import (
	"log"
	"server/config"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORM model for all Invoice records in the database
type Invoice struct {
	gorm.Model
	AppointmentID    uint   `gorm:"column:user_id" json:"user_id"`                     // ID of invoice that invoice is associated with
	OriginalBalance  int    `gorm:"column:original_balance" json:"original_balance"`   // Total original balance of the invoice (in cents)
	RemainingBalance int    `gorm:"column:remaining_balance" json:"remaining_balance"` // Remaining balance of the invoice (in cents)
	Status           string `gorm:"column:status" json:"status"`                       // Enforced list of statuses based on remaining balance (Unpaid, Paid, Overpaid)
}

/*
*Description*

func GetID

# Returns ID field from Invoice object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the invoice object
*/
func (invoice *Invoice) GetID() uint {
	return invoice.ID
}

/*
*Description*

func Create

Creates a new Invoice record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*Invoice>

		The created Invoice record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (invoice *Invoice) Create(db *gorm.DB) (map[string]Model, error) {
	err := db.Create(&invoice).Error
	returnRecords := map[string]Model{"invoice": invoice}
	return returnRecords, err
}

/*
*Description*

func Get

Retrieves a Invoice record in the database by ID if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	invoiceID  <uint>

		The ID of the invoice record being requested.

*Returns*

	_  <*Invoice>

		The Invoice record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (invoice *Invoice) Get(db *gorm.DB, invoiceID uint) (map[string]Model, error) {
	err := db.First(&invoice, invoiceID).Error
	returnRecords := map[string]Model{"invoice": invoice}
	return returnRecords, err
}

/*
*Description*

func GetAll

Retrieves all Invoice records from the database.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that the records will be retrieved from.

*Returns*

	_  <[]Invoice>

		The list of Invoice records that are retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (invoice *Invoice) GetAll(db *gorm.DB) ([]Invoice, error) {
	var invoices []Invoice
	err := db.Find(&invoices).Error

	return invoices, err
}

/*
*Description*

func Update

Updates the specified Invoice record in the database with the specified changes if the record exists.

Returns the updated record along with any errors that are thrown.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve and update the specified record.

	invoiceID  <uint>

		The ID of the invoice record being updated.

	updates  <map[string]interface{}>

		JSON with the fields that will be updated as keys and the updated values as values.

		Ex:
			{
				"name": "New name",
				"address": "New address"
			}

*Returns*

	_  <*Invoice>

		The Invoice record that is updated in the database.

	_  <error>

		Encountered error (nil if no errors are encountered)
*/
func (invoice *Invoice) Update(db *gorm.DB, invoiceID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm invoiceID exists in the database and get current object
	returnRecords, err := invoice.Get(db, invoiceID)
	updateInvoice := returnRecords["invoice"]

	if err != nil {
		return returnRecords, err
	}

	err = db.Model(&updateInvoice).Clauses(clause.Returning{}).Where("id = ?", invoiceID).Updates(updates).Error
	returnRecords = map[string]Model{"invoice": updateInvoice}

	return returnRecords, err
}

/*
*Description*

func Delete

Deletes the specified Invoice record from the database if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be deleted.

	invoiceID  <uint>

		The ID of the invoice record being deleted.

*Returns*

	_  <*Invoice>

		The deleted Invoice record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (invoice *Invoice) Delete(db *gorm.DB, invoiceID uint) (map[string]Model, error) {
	// Confirm invoiceID exists in the database and get current object
	returnRecords, err := invoice.Get(db, invoiceID)
	deleteInvoice := returnRecords["invoice"]

	if err != nil {
		return returnRecords, err
	}

	if config.Debug {
		log.Printf("\n\nInvoice object targeted for deletion:\n\n%+v\n\n", deleteInvoice)
	}

	err = db.Delete(deleteInvoice).Error
	returnRecords = map[string]Model{"invoice": deleteInvoice}

	return returnRecords, err
}
