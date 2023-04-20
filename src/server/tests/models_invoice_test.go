package tests

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestCreateGetInvoice

Tests the Create and Get methods for the Invoice db object. Confirms that the created Invoice object is returned when the method is called and that the record is created in the application database.
*/
func TestCreateGetInvoice(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create invoice record
	testCreateInvoice := models.Invoice{
		AppointmentID:    8675,
		OriginalBalance:  2000,
		RemainingBalance: 1000,
	}

	returnRecords, err := testCreateInvoice.Create(testAppDB)
	createdInvoice := returnRecords["invoice"]
	if err != nil {
		t.Errorf("Could not create test Invoice. %s", err)
	}

	// Attempt to retrieve newly created invoice record from database
	testGetInvoice := models.Invoice{}

	returnRecords, err = testGetInvoice.Get(testAppDB, createdInvoice.GetID())
	returnedInvoice := returnRecords["invoice"]
	if err != nil {
		t.Errorf("func Invoice.Get failed to return test Invoice. %s", err)
	}

	unequalFields, equal := models.Equal(createdInvoice, returnedInvoice)
	assert.Truef(t, equal, "The following fields did not match between the created and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestUpdateInvoice

Tests the Update method for the Invoice db object. Confirmed that the updated Invoice object is returned and that the record was updated in the datbas. Throws the appropriate error if the record doesn't exist in the database
*/
func TestUpdateInvoice(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Invoice record
	testInvoice := &models.Invoice{
		AppointmentID:    8675,
		OriginalBalance:  2000,
		RemainingBalance: 1000,
	}

	returnRecords, err := testInvoice.Create(testAppDB)
	createdInvoice := returnRecords["invoice"]
	if err != nil {
		t.Errorf("Could not create test Invoice.  --  %s", err)
	}

	invoiceID := createdInvoice.GetID()

	// Attempt to update created Invoice record
	testUpdateInvoice := models.Invoice{}

	updates := map[string]interface{}{
		"remaining_balance": 0,
	}

	returnRecords, err = testUpdateInvoice.Update(testAppDB, invoiceID, updates)
	updatedInvoice := returnRecords["invoice"]
	if err != nil {
		t.Errorf("Could not update test Invoice.  --  %s", err)
	}

	// Attempt to retrieve created/updated invoice record from database
	testGetInvoice := models.Invoice{}

	returnRecords, err = testGetInvoice.Get(testAppDB, invoiceID)
	returnedInvoice := returnRecords["invoice"]
	if err != nil {
		t.Errorf("Could not retrieve created/updated test Invoice record from teh database.  --  %s", err)
	}

	unequalFields, equal := models.Equal(updatedInvoice, returnedInvoice)
	assert.Truef(t, equal, "The following fields did not match between the updated and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestDeleteInvoice

Tests the Delete method for the Invoice db object. Confirms that the deleted Invoice object is returned when the method is called and that the record is deleted from the DB. Throws the appropriate error if the record doesn't exist.
*/
func TestDeleteInvoice(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Invoice record
	testInvoice := &models.Invoice{
		AppointmentID:    8675,
		OriginalBalance:  2000,
		RemainingBalance: 1000,
	}

	returnRecords, err := testInvoice.Create(testAppDB)
	createdInvoice := returnRecords["invoice"]
	if err != nil {
		t.Errorf("Could not create test Invoice.  --  %s", err)
	}

	invoiceID := createdInvoice.GetID()

	// Attempt to delete the Invoice record that was created in the database
	testDeleteInvoice := models.Invoice{}

	_, err = testDeleteInvoice.Delete(testAppDB, invoiceID)
	if err != nil {
		t.Errorf("Could not delete test Invoice.  --  %s", err)
	}

	// Check to see if deleted record's ID still exists in the database
	checkInvoice := models.Invoice{}

	deletedIDExists, err := checkInvoice.IDExists(testAppDB, invoiceID)
	if err != nil {
		t.Errorf("Could not confirm if deleted Invoice record's ID still exists in the database.")
	}
	assert.False(t, deletedIDExists, "Invoice record's ID still exists in the database after Delete method was executed.")
}
