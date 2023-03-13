package tests

import (
	"fmt"
	"server/database"
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
	database.FormatAllTables(testAppDB)

	testCreateUser := models.User{
		Email:       "test@gmail.com",
		Username:    "user1",
		Password:    "pw123",
		AccountType: "Individual",
		FirstName:   "John",
		LastName:    "Smith",
	}

	createdUser, err := testCreateUser.CreateUser(testAppDB)
	if err != nil {
		t.Fatalf("ERROR:  Could not create test User. %s", err)
	}
	createdUserID := fmt.Sprintf("%d", createdUser.ID)

	testGetUser := models.User{}
	returnedUser, err := testGetUser.GetUser(testAppDB, createdUserID)
	if err != nil {
		t.Fatalf("ERROR:  func GetUser failed to return test User. %s", err)
	}

	unequalFields, equal := createdUser.Equal(returnedUser)
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
