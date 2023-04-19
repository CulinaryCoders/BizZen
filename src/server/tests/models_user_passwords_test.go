package tests

import (
	"server/config"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestHashPassword

Tests the HashPassword method to ensure that user passwords are being hashed as expected.
*/
func TestHashPassword(t *testing.T) {
	var plainTextPassword string = "to8Od5DGg"
	var expectedHashString string = "$2a$08$IUhUwqZdaRG1K0.IRUqBj.X2zOZpBQWAAzHFKj6cp1kmIxKnLVSuq"

	returnedHashString, err := models.HashPassword(plainTextPassword, config.PWHashCost)

	if err != nil {
		t.Errorf("Error returned when one wasn't expected. ERROR:  %s", err)
	}

	assert.Equal(t, returnedHashString, expectedHashString, "Returned password should match expected hashed password.\n\nReturned password:  %s\nExpected hashed password:  %s\n\n", returnedHashString, expectedHashString)
}

/*
*Description*

func TestCheckPassword

Tests the CheckPassword method to ensure the method is correctly confirming when passwords match and throwing the appropriate error/response when they don't.
*/
func TestCheckPassword(t *testing.T) {
	testUserPassword := "Jzb!yxK@Ito5h&A_1"
	hashedPassword, err := models.HashPassword(testUserPassword, config.PWHashCost)

	if err != nil {
		t.Errorf("Error returned by HashPassword when one wasn't expected. ERROR:  %s", err)
	}

	user := &models.User{
		Password: hashedPassword,
	}

	//  Confirm correct password doesn't return an error
	err = user.CheckPassword(testUserPassword)
	if err != nil {
		t.Errorf("CASE [Correct Password]:  Error returned by CheckPassword when one wasn't expected. ERROR:  %s", err)
	}

	//  Confirm incorrect password returns an error
	err = user.CheckPassword("WrongPassword")
	if err == nil {
		t.Error("CASE [Incorrect Password]:  CheckPassword did not return error when incorrect password was provided.")
	}
}
