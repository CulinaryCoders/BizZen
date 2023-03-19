package sample_data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"server/models"

	"gorm.io/gorm"
)

// TODO:  Add documentation (type SampleData)
type SampleData struct {
	Users            []*models.User            `json:"users"`
	Businesses       []*models.Business        `json:"businesses"`
	Offices          []*models.Office          `json:"offices"`
	Addresses        []*models.Address         `json:"addresses"`
	Contacts         []*models.ContactInfo     `json:"contacts"`
	Services         []*models.Service         `json:"services"`
	ServiceOfferings []*models.ServiceOffering `json:"service_offerings"`
}

// TODO:  Add documentation (type DataLoadMapping)
type DataLoadMapping[Model models.Model] struct {
	Records                   []Model
	PrimaryReturnObjectKey    string
	SecondaryReturnObjectKeys []string
}

// TODO:  Add documentation (func LoadJSONSampleData)
func (dataLoadMapping DataLoadMapping[Model]) createSampleRecords(db *gorm.DB) error {
	err := createSampleRecords(db, dataLoadMapping.Records, dataLoadMapping.PrimaryReturnObjectKey, dataLoadMapping.SecondaryReturnObjectKeys...)
	return err
}

// TODO:  Add documentation (func LoadJSONSampleData)
func createSampleRecords[model models.Model](db *gorm.DB, records []model, primaryObjectKey string, secondaryReturnObjectKeys ...string) error {

	log.Printf("Number of '%s' records in JSON file:  %d", primaryObjectKey, len(records))

	for i := 0; i < len(records); i++ {

		returnRecords, err := records[i].Create(db)
		primaryRecord, primaryKeyExists := returnRecords[primaryObjectKey]
		if err != nil {
			return err
		} else if !primaryKeyExists {
			errorMessage := fmt.Sprintf("Primary key ('%s') does not exist in map returned from object creation.", primaryObjectKey)
			return errors.New(errorMessage)
		}

		log.Printf("Primary object created ('%s'):\n\n%+v\n\n", primaryObjectKey, primaryRecord)

		if len(secondaryReturnObjectKeys) > 0 {
			for j := 0; j < len(secondaryReturnObjectKeys); j++ {

				secondaryReturnObjectKey := secondaryReturnObjectKeys[j]
				secondaryRecord, secondaryKeyExists := returnRecords[secondaryReturnObjectKey]
				if !secondaryKeyExists {
					errorMessage := fmt.Sprintf("Secondary key ('%s') does not exist in map returned from object creation.", secondaryReturnObjectKey)
					return errors.New(errorMessage)
				}

				log.Printf("Secondary object created ('%s' from '%s' object creation):\n\n%+v\n\n", secondaryReturnObjectKey, primaryObjectKey, secondaryRecord)

			}
		}
	}

	return nil
}

// TODO:  Add documentation (func LoadJSONSampleData)
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
	userLoadMapping := DataLoadMapping[*models.User]{
		Records:                   sampleData.Users,
		PrimaryReturnObjectKey:    "user",
		SecondaryReturnObjectKeys: []string{},
	}
	userLoadMapping.createSampleRecords(db)
	//  Businesses / Main Offices
	businessLoadMapping := DataLoadMapping[*models.Business]{
		Records:                   sampleData.Businesses,
		PrimaryReturnObjectKey:    "business",
		SecondaryReturnObjectKeys: []string{"office"},
	}
	businessLoadMapping.createSampleRecords(db)
	//  Addresses
	addressLoadMapping := DataLoadMapping[*models.Address]{
		Records:                   sampleData.Addresses,
		PrimaryReturnObjectKey:    "address",
		SecondaryReturnObjectKeys: []string{},
	}
	addressLoadMapping.createSampleRecords(db)

	return nil
}
