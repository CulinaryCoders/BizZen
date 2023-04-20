package tests

import (
	"server/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestCreateGetService

Tests the Create and Get methods for the Service db object. Confirms that the created Service object is returned when the method is called and that the record is created in the application database.
*/
func TestCreateGetService(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create service record
	testCreateService := models.Service{
		BusinessID:    265,
		Name:          "Geriatric Aqua Arobics",
		Description:   "Reserved for elders who are truly bold and noble.",
		StartDateTime: time.Date(2023, 5, 4, 12, 00, 00, 00, time.Local),
		Length:        120,
		Capacity:      40,
		CancelFee:     0,
		Price:         1000,
	}

	returnRecords, err := testCreateService.Create(testAppDB)
	createdService := returnRecords["service"]
	if err != nil {
		t.Errorf("Could not create test Service. %s", err)
	}

	// Attempt to retrieve newly created service record from database
	testGetService := models.Service{}

	returnRecords, err = testGetService.Get(testAppDB, createdService.GetID())
	returnedService := returnRecords["service"]
	if err != nil {
		t.Errorf("func Service.Get failed to return test Service. %s", err)
	}

	unequalFields, equal := models.Equal(createdService, returnedService)
	assert.Truef(t, equal, "The following fields did not match between the created and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestUpdateService

Tests the Update method for the Service db object. Confirmed that the updated Service object is returned and that the record was updated in the datbas. Throws the appropriate error if the record doesn't exist in the database
*/
func TestUpdateService(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Service record
	testService := &models.Service{
		BusinessID:    128,
		Name:          "Planks & Pilates",
		Description:   "I've heard Kylo Ren has an 8-pack. That Kylo Ren is shredded.",
		StartDateTime: time.Date(2023, 04, 20, 17, 30, 00, 00, time.Local),
		Length:        30,
		Capacity:      20,
		CancelFee:     0,
		Price:         2000,
	}

	returnRecords, err := testService.Create(testAppDB)
	createdService := returnRecords["service"]
	if err != nil {
		t.Errorf("Could not create test Service.  --  %s", err)
	}

	serviceID := createdService.GetID()

	// Attempt to update created Service record
	testUpdateService := models.Service{}

	updates := map[string]interface{}{
		"name":   "Planks, Pilates, and Pushups",
		"length": 60,
		"price":  2500,
	}

	returnRecords, err = testUpdateService.Update(testAppDB, serviceID, updates)
	updatedService := returnRecords["service"]
	if err != nil {
		t.Errorf("Could not update test Service.  --  %s", err)
	}

	// Attempt to retrieve created/updated service record from database
	testGetService := models.Service{}

	returnRecords, err = testGetService.Get(testAppDB, serviceID)
	returnedService := returnRecords["service"]
	if err != nil {
		t.Errorf("Could not retrieve created/updated test Service record from teh database.  --  %s", err)
	}

	unequalFields, equal := models.Equal(updatedService, returnedService)
	assert.Truef(t, equal, "The following fields did not match between the updated and returned object  --  %s", unequalFields)
}

/*
*Description*

func TestDeleteService

Tests the Delete method for the Service db object. Confirms that the deleted Service object is returned when the method is called and that the record is deleted from the DB. Throws the appropriate error if the record doesn't exist.
*/
func TestDeleteService(t *testing.T) {
	// Refresh database to control testing environment
	models.FormatAllTables(testAppDB)

	// Create Service record
	testService := &models.Service{
		BusinessID:    99,
		Name:          "Spirited Spin Class",
		Description:   "Spin class with spirit!",
		StartDateTime: time.Date(2023, 11, 5, 10, 30, 00, 00, time.Local),
		Length:        30,
		Capacity:      12,
		CancelFee:     0,
		Price:         3000,
	}

	returnRecords, err := testService.Create(testAppDB)
	createdService := returnRecords["service"]
	if err != nil {
		t.Errorf("Could not create test Service.  --  %s", err)
	}

	serviceID := createdService.GetID()

	// Attempt to delete the Service record that was created in the database
	testDeleteService := models.Service{}

	_, err = testDeleteService.Delete(testAppDB, serviceID)
	if err != nil {
		t.Errorf("Could not delete test Service.  --  %s", err)
	}

	// Check to see if deleted record's ID still exists in the database
	checkService := models.Service{}

	deletedIDExists, err := checkService.IDExists(testAppDB, serviceID)
	if err != nil {
		t.Errorf("Could not confirm if deleted Service record's ID still exists in the database.")
	}
	assert.False(t, deletedIDExists, "Service record's ID still exists in the database after Delete method was executed.")
}
