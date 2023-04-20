package tests

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestStandardizeUserFields

Tests the User.StandardizeFields method to confirm that fields are being standardized as expected.
*/
func TestStandardizeUserFields(t *testing.T) {

	// Test user
	testUser := models.User{
		Email:       "TeST-oneTWOoneTWO@yahoo.com  ",
		Password:    "pw123",
		AccountType: "  useR  ",
		FirstName:   "  jessiE",
		LastName:    " PinKMAn ",
	}

	// Define expected values
	expectedEmail := "test-onetwoonetwo@yahoo.com"
	expectedPassword := "pw123"
	expectedAccountType := "User"
	expectedFirstName := "Jessie"
	expectedLastName := "Pinkman"

	// Standardize fields
	testUser.StandardizeFields()

	// Confirm standardized fields match expected values
	assert.Equal(t, testUser.Email, expectedEmail, "Standardized user email (%s) should match expected email (%s).", testUser.Email, expectedEmail)
	assert.Equal(t, testUser.Password, expectedPassword, "Standardized password (%s) should match expected password (%s).", testUser.Password, expectedPassword)
	assert.Equal(t, testUser.AccountType, expectedAccountType, "Standardized account type (%s) should match expected account type (%s).", testUser.AccountType, expectedAccountType)
	assert.Equal(t, testUser.FirstName, expectedFirstName, "Standardized first name (%s) should match expected first name (%s).", testUser.FirstName, expectedFirstName)
	assert.Equal(t, testUser.LastName, expectedLastName, "Standardized last name (%s) should match expected last name (%s).", testUser.LastName, expectedLastName)
}

/*
*Description*

func TestCreateGetUser

Tests the Create and Get methods for the User db object. Confirms that the created User object is returned when the method is called and that the record is created in the application  database.
*/
func TestCreateGetUser(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create user record
	testCreateUser := &models.User{
		Email:       "test@gmail.com",
		Password:    "pw123",
		AccountType: "User",
		FirstName:   "John",
		LastName:    "Smith",
	}

	returnRecords, err := testCreateUser.Create(testAppDB)
	createdUser := returnRecords["user"]
	if err != nil {
		t.Errorf("Could not create test User.  --  %s", err)
	}

	// Attempt to retrieve newly created user record from database
	testGetUser := models.User{}

	returnRecords, err = testGetUser.Get(testAppDB, createdUser.GetID())
	returnedUser := returnRecords["user"]
	if err != nil {
		t.Errorf("func User.Get failed to return test User  --  %s", err)
	}

	unequalFields, equal := models.Equal(createdUser, returnedUser)
	assert.Truef(t, equal, "The following fields did not match between the created and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestUpdateUser

Tests the Update method for the User db object. Confirmed that the updated User object is returned and that the record was updated in the datbas. Throws the appropriate error if the record doesn't exist in the database
*/
func TestUpdateUser(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create user record
	testUser := &models.User{
		Email:       "MakingMyWayDowntown123@hotmail.com",
		Password:    "Walking_fast",
		AccountType: "Business",
		FirstName:   "Vanessa",
		LastName:    "Carlton",
	}

	returnRecords, err := testUser.Create(testAppDB)
	createdUser := returnRecords["user"]
	if err != nil {
		t.Errorf("Could not create test User.  --  %s", err)
	}

	userID := createdUser.GetID()

	// Attempt to update created User record
	testUpdateUser := models.User{}

	updates := map[string]interface{}{
		"email":      "AndNowIWonder90210@uphoenix.edu",
		"first_name": "Definitely not Vanessa",
	}

	returnRecords, err = testUpdateUser.Update(testAppDB, userID, updates)
	updatedUser := returnRecords["user"]
	if err != nil {
		t.Errorf("Could not update test User.  --  %s", err)
	}

	// Attempt to retrieve created/updated user record from database
	testGetUser := models.User{}

	returnRecords, err = testGetUser.Get(testAppDB, userID)
	returnedUser := returnRecords["user"]
	if err != nil {
		t.Errorf("Could not retrieve created/updated test User record from teh database.  --  %s", err)
	}

	unequalFields, equal := models.Equal(updatedUser, returnedUser)
	assert.Truef(t, equal, "The following fields did not match between the updated and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestDeleteUser

Tests the Delete method for the User db object. Confirms that the deleted User object is returned when the method is called and that the record is deleted from the DB. Throws the appropriate error if the record doesn't exist.
*/
func TestDeleteUser(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create user record
	testUser := &models.User{
		Email:       "MakingMyWayDowntown123@hotmail.com",
		Password:    "Walking_fast",
		AccountType: "Business",
		FirstName:   "Vanessa",
		LastName:    "Carlton",
	}

	returnRecords, err := testUser.Create(testAppDB)
	createdUser := returnRecords["user"]
	if err != nil {
		t.Errorf("Could not create test User.  --  %s", err)
	}

	userID := createdUser.GetID()

	// Attempt to delete the User record that was created in the database
	testDeleteUser := models.User{}

	_, err = testDeleteUser.Delete(testAppDB, userID)
	if err != nil {
		t.Errorf("Could not delete test User.  --  %s", err)
	}

	// Check to see if deleted record's ID still exists in the database
	checkUser := models.User{}

	deletedIDExists, err := checkUser.IDExists(testAppDB, userID)
	if err != nil {
		t.Errorf("Could not confirm if deleted User record's ID still exists in the database.")
	}
	assert.False(t, deletedIDExists, "User record's ID still exists in the database after Delete method was executed.")
}
