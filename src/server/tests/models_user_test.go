package tests

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO:  Add documentation (func TestAuthenticate)
func TestAuthenticate(t *testing.T) {
	assert.True(t, true)
}

// TODO:  Add documentation (func TestCreateUser)
func TestCreateUser(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create user record
	testCreateUser := models.User{
		Email:       "test@gmail.com",
		Password:    "pw123",
		AccountType: "Individual",
		FirstName:   "John",
		LastName:    "Smith",
	}

	returnRecords, err := testCreateUser.Create(testAppDB)
	createdUser := returnRecords["user"]
	if err != nil {
		t.Fatalf("ERROR:  Could not create test User. %s", err)
	}

	// Attempt to retrieve newly created user record from database
	testGetUser := models.User{}

	returnRecords, err = testGetUser.Get(testAppDB, createdUser.GetID())
	returnedUser := returnRecords["user"]
	if err != nil {
		t.Fatalf("ERROR:  func GetUser failed to return test User. %s", err)
	}

	unequalFields, equal := models.Equal(createdUser, returnedUser)
	assert.Truef(t, equal, "ERROR: The following fields did not match between the created and returned object  --  %s", unequalFields)
}

// TODO:  Add documentation (func TestGetUser)
func TestGetUser(t *testing.T) {
	assert.True(t, true)
}

// TODO:  Add documentation (func TestUpdateUser)
func TestUpdateUser(t *testing.T) {
	assert.True(t, true)
}

// TODO:  Add documentation (func TestDeleteUser)
func TestDeleteUser(t *testing.T) {
	assert.True(t, true)
}
