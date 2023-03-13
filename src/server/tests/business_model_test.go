package tests

import (
	"fmt"
	"server/database"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO:  Add documentation (func TestCreateBusiness)
func TestCreateBusiness(t *testing.T) {
	// Refresh database to control testing environment
	database.FormatAllTables(testAppDB)

	testCreateBusiness := models.Business{
		OwnerID:      1,
		MainOfficeID: 1,
		Name:         "Gator Aider LLC",
		Type:         "Tutoring",
	}

	createdBusiness, err := testCreateBusiness.CreateBusiness(testAppDB)
	if err != nil {
		t.Fatalf("ERROR:  Could not create test Business. %s", err)
	}
	createdBusinessID := fmt.Sprintf("%d", createdBusiness.ID)

	testGetBusiness := models.Business{}
	returnedBusiness, err := testGetBusiness.GetBusiness(testAppDB, createdBusinessID)
	if err != nil {
		t.Fatalf("ERROR:  func GetBusiness failed to return test Business. %s", err)
	}

	unequalFields, equal := createdBusiness.Equal(returnedBusiness)
	assert.Truef(t, equal, "ERROR: The following fields did not match between the created and returned object  --  %s", unequalFields)
}

// TODO:  Add documentation (func TestGetBusiness)
func TestGetBusiness(t *testing.T) {
	assert.True(t, true)
}

// TODO:  Add documentation (func TestUpdateBusiness)
func TestUpdateBusiness(t *testing.T) {
	assert.True(t, true)
}

// TODO:  Add documentation (func TestDeleteBusiness)
func TestDeleteBusiness(t *testing.T) {
	assert.True(t, true)
}
