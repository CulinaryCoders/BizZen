package tests

import (
	"server/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestCreateGetAppointment

Tests the Create and Get methods for the Appointment db object. Confirms that the created Appointment object is returned when the method is called and that the record is created in the application database.
*/
func TestCreateGetAppointment(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create appointment record
	testCreateAppointment := models.Appointment{
		UserID:    69,
		ServiceID: 420,
	}

	returnRecords, err := testCreateAppointment.Create(testAppDB)
	createdAppointment := returnRecords["appointment"]
	if err != nil {
		t.Errorf("Could not create test Appointment. %s", err)
	}

	// Attempt to retrieve newly created Appointment record from database
	testGetAppointment := models.Appointment{}

	returnRecords, err = testGetAppointment.Get(testAppDB, createdAppointment.GetID())
	returnedAppointment := returnRecords["appointment"]
	if err != nil {
		t.Errorf("func Appointment.Get failed to return test Appointment. %s", err)
	}

	unequalFields, equal := models.Equal(createdAppointment, returnedAppointment)
	assert.Truef(t, equal, "The following fields did not match between the created and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestUpdateAppointment

Tests the Update method for the Appointment db object. Confirmed that the updated Appointment object is returned and that the record was updated in the datbas. Throws the appropriate error if the record doesn't exist in the database
*/
func TestUpdateAppointment(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Appointment record
	testAppointment := models.Appointment{
		UserID:    69,
		ServiceID: 420,
	}

	returnRecords, err := testAppointment.Create(testAppDB)
	createdAppointment := returnRecords["appointment"]
	if err != nil {
		t.Errorf("Could not create test Appointment.  --  %s", err)
	}

	apptID := createdAppointment.GetID()

	// Attempt to update created Appointment record
	testUpdateAppointment := models.Appointment{}

	updates := map[string]interface{}{
		"active":           "false",
		"cancel_date_time": time.Now(),
	}

	returnRecords, err = testUpdateAppointment.Update(testAppDB, apptID, updates)
	updatedAppointment := returnRecords["appointment"]
	if err != nil {
		t.Errorf("Could not update test Appointment.  --  %s", err)
	}

	// Attempt to retrieve created/updated Appointment record from database
	testGetAppointment := models.Appointment{}

	returnRecords, err = testGetAppointment.Get(testAppDB, apptID)
	returnedAppointment := returnRecords["appointment"]
	if err != nil {
		t.Errorf("Could not retrieve created/updated test Appointment record from teh database.  --  %s", err)
	}

	unequalFields, equal := models.Equal(updatedAppointment, returnedAppointment)
	assert.Truef(t, equal, "The following fields did not match between the updated and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestDeleteAppointment

Tests the Delete method for the Appointment db object. Confirms that the deleted Appointment object is returned when the method is called and that the record is deleted from the DB. Throws the appropriate error if the record doesn't exist.
*/
func TestDeleteAppointment(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Appointment record
	testAppointment := &models.Appointment{
		UserID:    69,
		ServiceID: 420,
	}

	returnRecords, err := testAppointment.Create(testAppDB)
	createdAppointment := returnRecords["appointment"]
	if err != nil {
		t.Errorf("Could not create test Appointment.  --  %s", err)
	}

	apptID := createdAppointment.GetID()

	// Attempt to delete the Appointment record that was created in the database
	testDeleteAppointment := models.Appointment{}

	_, err = testDeleteAppointment.Delete(testAppDB, apptID)
	if err != nil {
		t.Errorf("Could not delete test Appointment.  --  %s", err)
	}

	// Check to see if deleted record's ID still exists in the database
	checkAppointment := models.Appointment{}

	deletedIDExists, err := checkAppointment.IDExists(testAppDB, apptID)
	if err != nil {
		t.Errorf("Could not confirm if deleted Appointment record's ID still exists in the database.")
	}
	assert.False(t, deletedIDExists, "Appointment record's ID still exists in the database after Delete method was executed.")
}
