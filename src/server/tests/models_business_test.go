package tests

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestCreateGetBusiness

Tests the Create and Get methods for the Business db object. Confirms that the created Business object is returned when the method is called and that the record is created in the application database.
*/
func TestCreateGetBusiness(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create business record
	testCreateBusiness := models.Business{
		OwnerID: 1,
		Name:    "Gator Aider LLC",
	}

	returnRecords, err := testCreateBusiness.Create(testAppDB)
	createdBusiness := returnRecords["business"]
	if err != nil {
		t.Errorf("Could not create test Business. %s", err)
	}

	// Attempt to retrieve newly created business record from database
	testGetBusiness := models.Business{}

	returnRecords, err = testGetBusiness.Get(testAppDB, createdBusiness.GetID())
	returnedBusiness := returnRecords["business"]
	if err != nil {
		t.Errorf("func Business.Get failed to return test Business. %s", err)
	}

	unequalFields, equal := models.Equal(createdBusiness, returnedBusiness)
	assert.Truef(t, equal, "The following fields did not match between the created and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestUpdateBusiness

Tests the Update method for the Business db object. Confirmed that the updated Business object is returned and that the record was updated in the datbas. Throws the appropriate error if the record doesn't exist in the database
*/
func TestUpdateBusiness(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Business record
	testBusiness := &models.Business{
		OwnerID: 74,
		Name:    "TOP SECRET - Government Alien Technologies",
	}

	returnRecords, err := testBusiness.Create(testAppDB)
	createdBusiness := returnRecords["business"]
	if err != nil {
		t.Errorf("Could not create test Business.  --  %s", err)
	}

	businessID := createdBusiness.GetID()

	// Attempt to update created Business record
	testUpdateBusiness := models.Business{}

	updates := map[string]interface{}{
		"name": "Nothing to See Here, Inc.",
	}

	returnRecords, err = testUpdateBusiness.Update(testAppDB, businessID, updates)
	updatedBusiness := returnRecords["business"]
	if err != nil {
		t.Errorf("Could not update test Business.  --  %s", err)
	}

	// Attempt to retrieve created/updated business record from database
	testGetBusiness := models.Business{}

	returnRecords, err = testGetBusiness.Get(testAppDB, businessID)
	returnedBusiness := returnRecords["business"]
	if err != nil {
		t.Errorf("Could not retrieve created/updated test Business record from teh database.  --  %s", err)
	}

	unequalFields, equal := models.Equal(updatedBusiness, returnedBusiness)
	assert.Truef(t, equal, "The following fields did not match between the updated and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestDeleteBusiness

Tests the Delete method for the Business db object. Confirms that the deleted Business object is returned when the method is called and that the record is deleted from the DB. Throws the appropriate error if the record doesn't exist.
*/
func TestDeleteBusiness(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Business record
	testBusiness := &models.Business{
		OwnerID: 58,
		Name:    "Michael Meyers Reform School",
	}

	returnRecords, err := testBusiness.Create(testAppDB)
	createdBusiness := returnRecords["business"]
	if err != nil {
		t.Errorf("Could not create test Business.  --  %s", err)
	}

	businessID := createdBusiness.GetID()

	// Attempt to delete the Business record that was created in the database
	testDeleteBusiness := models.Business{}

	_, err = testDeleteBusiness.Delete(testAppDB, businessID)
	if err != nil {
		t.Errorf("Could not delete test Business.  --  %s", err)
	}

	// Check to see if deleted record's ID still exists in the database
	checkBusiness := models.Business{}

	deletedIDExists, err := checkBusiness.IDExists(testAppDB, businessID)
	if err != nil {
		t.Errorf("Could not confirm if deleted Business record's ID still exists in the database.")
	}
	assert.False(t, deletedIDExists, "Business record's ID still exists in the database after Delete method was executed.")
}
