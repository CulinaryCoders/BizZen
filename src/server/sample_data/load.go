package sample_data

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"server/models"

	"gorm.io/gorm"
)

type SampleData struct {
	Users            []models.User            `json:"users"`
	Businesses       []models.Business        `json:"businesses"`
	Offices          []models.Office          `json:"offices"`
	Addresses        []*models.Address        `json:"addresses"`
	Contacts         []models.ContactInfo     `json:"contacts"`
	Services         []models.Service         `json:"services"`
	ServiceOfferings []models.ServiceOffering `json:"service_offerings"`
}

type ModelsMap struct {
	Models map[int]models.Model
}

func createSampleRecords[model models.Model](db *gorm.DB, records []model, objectName string) error {
	// TODO: Create generic interface for all models (implement Create, Update, Get, Delete, Equal) and incorporate into test load
	log.Printf("Number of %s records in JSON file:  %d", objectName, len(records))
	for i := 0; i < len(records); i++ {
		createdAddress, err := records[i].Create(db)
		if err != nil {
			return err
		}

		log.Println(createdAddress)
	}

	return nil
}

func LoadJSONSampleData(db *gorm.DB) error {

	serverExe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	serverExePath := filepath.Dir(serverExe)
	sampleDataBasePath := filepath.Join(serverExePath, "sample_data")

	jsonFileName := "sample-data.json"
	jsonFilePath := filepath.Join(sampleDataBasePath, jsonFileName)

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

	// convertModelSliceToMap := func(modelSlice []any) map[int]models.Model {
	// 	var newMap map[int]models.Model

	// 	for i := 0; i < len(modelSlice); i++ {
	// 		modelInterface := reflect.TypeOf(new(models.Model)).Elem()
	// 		if reflect.TypeOf(modelSlice[i]).Implements(modelInterface) {
	// 			newMap[i+1] = modelSlice[i]
	// 		}
	// 	}

	// 	return newMap
	// }

	// convertedAddressMap := convertModelSliceToMap(sampleData.Addresses)
	// var addressesMap ModelsMap = ModelsMap{Models: sampleData.Addresses}

	err = createSampleRecords(db, sampleData.Addresses, "Address")
	if err != nil {
		return err
	}

	return nil
}
