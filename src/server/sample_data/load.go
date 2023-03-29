package sample_data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"server/config"
	"server/models"

	"gorm.io/gorm"
)

var debug bool = config.AppConfig.DEBUG_MODE

/*
SampleData type is used to store a list of object instances for each DB object type in the database.

Records for each object type can be created from a JSON file and then loaded into the appropriate DB instance.
*/
type SampleData struct {
	Users            []*models.User            `json:"users"`             // List of User records
	Businesses       []*models.Business        `json:"businesses"`        // List of Business records
	Offices          []*models.Office          `json:"offices"`           // List of Office records
	Addresses        []*models.Address         `json:"addresses"`         // List of Address records
	Contacts         []*models.ContactInfo     `json:"contacts"`          // List of Contact records
	Services         []*models.Service         `json:"services"`          // List of Service records
	ServiceOfferings []*models.ServiceOffering `json:"service_offerings"` // List of ServiceOffering records
}

// DataLoadMapping type is a generic wrapper struct designed to simplify the creation of records for all GORM DB object types that implement the 'Model' interface.
type DataLoadMapping[Model models.Model] struct {
	Records                   []Model  // List of records to be created
	PrimaryReturnObjectKey    string   // The JSON key that is returned for the primary object that is created in the database and returned by the 'Create' function call for that specific object type.
	SecondaryReturnObjectKeys []string // A list of JSON keys that are returned for any secondary objects that are created in the database and returned by the 'Create' function call for that specific object type.
}

/*
*Description*

func (dataLoadMapping DataLoadMapping[Model]) createSampleRecords

A generic wrapper function for the DataLoadMapping interface to be able to call the 'createSampleRecords' function for any
GORM DB object type that implements the 'Model' interface.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (dataLoadMapping DataLoadMapping[Model]) createSampleRecords(db *gorm.DB) error {
	err := createSampleRecords(db, dataLoadMapping.Records, dataLoadMapping.PrimaryReturnObjectKey, dataLoadMapping.SecondaryReturnObjectKeys...)
	return err
}

/*
*Description*

func createSampleRecords

Creates the list of records that are passed in the specified database instance and logs all of the objects that are created
and/or errors that are encountered while records are being created in the database.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

	records  <[]model>

		An array of database object instances whose type implements the generic 'Model' interface.

	primaryObjectKey  <string>

		The JSON key that is returned for the primary object that is created in the database and returned by the 'Create' function call for that specific object type.

	secondaryReturnObjectKeys  <[]string>

		An array of JSON keys that are returned for any secondary objects that are created in the database and returned by the 'Create' function call for that specific object type.

*Returns*

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func createSampleRecords[model models.Model](db *gorm.DB, records []model, primaryObjectKey string, secondaryReturnObjectKeys ...string) error {

	var recordCount int = len(records)
	var secondaryKeyCount int = len(secondaryReturnObjectKeys)

	if debug {
		log.Printf("Number of '%s' records in JSON file:  %d", primaryObjectKey, len(records))
	}

	for i := 0; i < recordCount; i++ {

		returnRecords, err := records[i].Create(db)
		primaryRecord, primaryKeyExists := returnRecords[primaryObjectKey]
		if err != nil {
			return err
		} else if !primaryKeyExists {
			errorMessage := fmt.Sprintf("Primary key ('%s') does not exist in map returned from object creation.", primaryObjectKey)
			return errors.New(errorMessage)
		}

		if debug {
			log.Printf("Primary object created ('%s'):\n\n%+v\n\n", primaryObjectKey, primaryRecord)
		}

		if secondaryKeyCount > 0 {
			for j := 0; j < secondaryKeyCount; j++ {

				secondaryReturnObjectKey := secondaryReturnObjectKeys[j]
				secondaryRecord, secondaryKeyExists := returnRecords[secondaryReturnObjectKey]
				if !secondaryKeyExists {
					errorMessage := fmt.Sprintf("Secondary key ('%s') does not exist in map returned from object creation.", secondaryReturnObjectKey)
					return errors.New(errorMessage)
				}

				if debug {
					log.Printf("Secondary object created ('%s' from '%s' object creation):\n\n%+v\n\n", secondaryReturnObjectKey, primaryObjectKey, secondaryRecord)
				}

			}
		}
	}

	var recordKeysLogString string = fmt.Sprintf("'%s'", primaryObjectKey)
	if secondaryKeyCount > 0 {
		for i := 0; i < secondaryKeyCount; i++ {
			recordKeysLogString = fmt.Sprintf("%s, '%s'", recordKeysLogString, secondaryReturnObjectKeys[i])
		}
	}

	log.Printf("%d record(s) successfully loaded from JSON file. (DB object types created: %s)", recordCount, recordKeysLogString)

	return nil
}

/*
*Description*

func LoadJSONSampleData

Loads all records within the 'sample-data.json' file into the appropriate GORM database object instance
and creates those records within the specified database instance.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func LoadJSONSampleData(db *gorm.DB) error {

	//  Get path of server.exe and set base path for JSON file that contains sample records for each object type
	serverExe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	serverExePath := filepath.Dir(serverExe)
	sampleDataFileBasePath := filepath.Join(serverExePath, "sample_data")

	jsonFileName := "sample-data.json"
	jsonFilePath := filepath.Join(sampleDataFileBasePath, jsonFileName)

	//  Read in/load data from JSON file that contains sample records for each object type
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	sampleDataBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var sampleData SampleData
	json.Unmarshal(sampleDataBytes, &sampleData)

	//  Create DataLoadMapping object for each DB object type and load sample records that are available for each
	//  Users
	var userJSONKey string = "user"

	userLoadMapping := DataLoadMapping[*models.User]{
		Records:                   sampleData.Users,
		PrimaryReturnObjectKey:    userJSONKey,
		SecondaryReturnObjectKeys: []string{},
	}

	err = userLoadMapping.createSampleRecords(db)
	if err != nil {
		return err
	}

	//  Businesses / Main Offices
	var businessJSONKey string = "business"
	var officeJSONKey string = "office"

	businessLoadMapping := DataLoadMapping[*models.Business]{
		Records:                   sampleData.Businesses,
		PrimaryReturnObjectKey:    businessJSONKey,
		SecondaryReturnObjectKeys: []string{officeJSONKey},
	}

	err = businessLoadMapping.createSampleRecords(db)
	if err != nil {
		return err
	}

	//  Addresses
	var addressJSONKey string = "address"

	addressLoadMapping := DataLoadMapping[*models.Address]{
		Records:                   sampleData.Addresses,
		PrimaryReturnObjectKey:    addressJSONKey,
		SecondaryReturnObjectKeys: []string{},
	}

	err = addressLoadMapping.createSampleRecords(db)
	if err != nil {
		return err
	}

	//  Services
	var serviceJSONKey string = "service"

	serviceLoadMapping := DataLoadMapping[*models.Service]{
		Records:                   sampleData.Services,
		PrimaryReturnObjectKey:    serviceJSONKey,
		SecondaryReturnObjectKeys: []string{},
	}

	err = serviceLoadMapping.createSampleRecords(db)
	if err != nil {
		return err
	}

	//  ServiceOfferings
	var serviceOfferingJSONKey string = "service_offering"

	serviceOfferingLoadMapping := DataLoadMapping[*models.ServiceOffering]{
		Records:                   sampleData.ServiceOfferings,
		PrimaryReturnObjectKey:    serviceOfferingJSONKey,
		SecondaryReturnObjectKeys: []string{},
	}

	err = serviceOfferingLoadMapping.createSampleRecords(db)
	if err != nil {
		return err
	}

	return nil
}
