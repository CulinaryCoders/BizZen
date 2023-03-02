package tests

import (
	"encoding/json"
	"fmt"
	"server/database"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var standardAddresses = []struct {
	Address models.Address
}{
	{Address: models.Address{
		Address1: "1234 Gator Way",
		Address2: "",
		City:     "Gainesville",
		State:    "FL",
		ZipCode:  "12345",
	}},
	{Address: models.Address{
		Address1: "145 Homer Simpson Ln",
		Address2: "Apt #5",
		City:     "Springfield",
		State:    "IN",
		ZipCode:  "09090",
	}},
	{Address: models.Address{
		Address1: "888 Cueball Street",
		Address2: "P.O. Box 97",
		City:     "Boston",
		State:    "MA",
		ZipCode:  "33445",
	}},
}
var missingRequiredFieldsAddresses []models.Address
var invalidFieldValuesAddresses []models.Address

// TODO:  Add documentation (func createTestAddressRecords)
func createTestAddressRecords([]models.Address) error {
	return nil
}

// TODO:  Add documentation (func TestCreateAddress)
func TestCreateAddress(t *testing.T) {
	// Refresh database to control testing environment
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
		t.Errorf("ERROR:  Could not create test address. %s", err)
	}
	createdAddressID := fmt.Sprintf("%d", createdAddress.ID)

	testGetAddress := models.Address{}
	returnedAddress, err := testGetAddress.GetAddress(testAppDB, createdAddressID)
	if err != nil {
		t.Errorf("ERROR:  func GetAddress failed to return test address. %s", err)
	}

	unequalFields, equal := createdAddress.Equal(returnedAddress)
	assert.Truef(t, equal, "ERROR: The following fields did not match between the created and returned object  --  %s", unequalFields)
}

// TODO:  Add documentation (func TestGetAddress)
func TestGetAddress(t *testing.T) {
	// Refresh database to control testing environment
	database.FormatAllTables(testAppDB)

	// Defined test cases
	// testCreateAddress1 := models.Address{
	// 	Address1: "1234 Gator Way",
	// 	Address2: "",
	// 	City:     "Gainesville",
	// 	State:    "FL",
	// 	ZipCode:  "12345",
	// }

	// testCreateAddress2 := models.Address{
	// 	Address1: "145 Homer Simpson Ln",
	// 	Address2: "Apt #5",
	// 	City:     "Springfield",
	// 	State:    "IN",
	// 	ZipCode:  "09090",
	// }

	// testCreateAddress3 := models.Address{
	// 	Address1: "888 Cueball Street",
	// 	Address2: "P.O. Box 97",
	// 	City:     "Boston",
	// 	State:    "MA",
	// 	ZipCode:  "33445",
	// }

	// Create list of test cases with expected outputs and/or errors
	type AddressTest struct {
		input         models.Address
		scenario      string
		expectedError error
	}

	var addressTests []AddressTest
	for _, standardCase := range standardAddresses {
		testDef := AddressTest{
			input:         standardCase.Address,
			scenario:      "Standard address (basic case)",
			expectedError: nil,
		}
		addressTests = append(addressTests, testDef)
	}

	// Iterate through test cases
	// TODO:  Add corner cases and test error handling (func TestGetAddress)
	for _, testCase := range addressTests {
		address := models.Address{}

		createdAddress, err := testCase.input.CreateAddress(testAppDB)
		if err != nil {
			t.Errorf("ERROR:  Could not create test address. %s", err)
		}
		createdAddressID := fmt.Sprintf("%d", createdAddress.ID)

		returnedAddress, err := address.GetAddress(testAppDB, createdAddressID)
		if err != nil {
			t.Errorf("ERROR:  func GetAddress failed to return test address. %s", err)
		}

		unequalFields, equal := returnedAddress.Equal(createdAddress)
		returnedAddressJSON, _ := json.Marshal(returnedAddress)
		createdAddressJSON, _ := json.Marshal(createdAddress)

		assert.Truef(t, equal, "ERROR: The following fields did not match between the created and returned object  --  %s\n\nOriginal Address (CreateAddress):  %v\n\nReturned Address (GetAddress):  %v", unequalFields, createdAddressJSON, returnedAddressJSON)
	}
}

// TODO:  Add documentation (func TestUpdateAddress)
func TestUpdateAddress(t *testing.T) {
	assert.True(t, true)
}

// TODO:  Add documentation (func TestDeleteAddress)
func TestDeleteAddress(t *testing.T) {
	assert.True(t, true)
}
