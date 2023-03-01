package tests

import (
	"fmt"
	"server/database"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAddress(t *testing.T) {
	database.FormatAllTables(testAppDB)

	testCreateAddress := models.Address{
		Address1: "1234 Gator Way",
		Address2: "",
		City:     "Gainesville",
		State:    "FL",
		ZipCode:  "12345",
	}

	createdAddress, err := testCreateAddress.CreateAddress(testAppDB)
	if err != nil {
		t.Fatalf("ERROR:  Could not create test address. %s", err)
	}
	createdAddressID := fmt.Sprintf("%d", createdAddress.ID)

	testGetAddress := models.Address{}
	returnedAddress, err := testGetAddress.GetAddress(testAppDB, createdAddressID)
	if err != nil {
		t.Fatalf("ERROR:  func GetAddress failed to return test address. %s", err)
	}

	unequalFields, equal := createdAddress.Equal(returnedAddress)
	assert.Truef(t, equal, "ERROR: The following fields did not match between the created and returned object  --  %s", unequalFields)
}
