package models

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

/*
Model is a generic interface that all GORM model database types implement.

This generic interface is intended to reduce the need for duplicative code and logic across this application, since every model implements these core methods.
*/
type Model interface {
	Create(db *gorm.DB) (map[string]Model, error)
	Get(db *gorm.DB, id uint) (map[string]Model, error)
	Update(db *gorm.DB, id uint, updates map[string]interface{}) (map[string]Model, error)
	Delete(db *gorm.DB, id uint) (map[string]Model, error)
	GetID() uint
	// TODO: IDExists
	// TODO: StandardizeFields
	// TODO: GetAttribute
	// TODO: GetAll
	// TODO: GetRecordsByPrimaryIDs
	// TODO: GetRecordsBySecondaryID
	// TODO: GetRecordsBySecondaryIDs
	// TODO: DeleteRecordsByPrimaryIDs
	// TODO: DeleteRecordsBySecondaryID
	// TODO: DeleteRecordsBySecondaryIDs
}

/*
*Description*

func Equal

Determines if two different objects that implement the Model interface are equal to each other (i.e. all fields match).

The primary purpose of this function is to test the functionality of database and handler calls to ensure that
the correct objects are being returned and/or updated in the database.

*Parameters*

	firstRecord  <Model>

		The first object/record to compare.

	secondRecord  <Model>

		The second object/record to compare.

*Returns*

	unequalFields  <[]string>

		The list of fields that did not match between the two objects being compared

	equal  <bool>

		If all the fields between the two objects are the same, true is returned. Otherwise, false is returned.
*/
func Equal(firstRecord Model, secondRecord Model) (unequalFields []string, equal bool) {
	equal = true

	//  Confirm that objects being compared are of the same type
	type1 := reflect.TypeOf(firstRecord).Elem()
	type2 := reflect.TypeOf(secondRecord).Elem()

	if type1 != type2 {
		equal = false

		unequalReason := fmt.Sprintf("Mismatched object types -- 'firstRecord' has a type of '%s' and 'secondRecord' has a type of '%s'", type1, type2)
		unequalFields = append(unequalFields, unequalReason)

		//  Don't perform any more checks if comparison objects have mismatched types to avoid additional errors
		return unequalFields, equal
	}

	// Compare IDs of two objects
	firstID := firstRecord.GetID()
	secondID := secondRecord.GetID()

	if firstID != secondID {
		unequalFields = append(unequalFields, "ID")
	}

	//  Generic object info
	objectType := type1 // Type assignment is arbitrary since objects being compared are confirmed to be of same type
	fieldCount := objectType.NumField()
	//  Specific field values for each object
	firstRecordValues := reflect.ValueOf(firstRecord).Elem()
	secondRecordValues := reflect.ValueOf(secondRecord).Elem()

	for i := 0; i < fieldCount; i++ {
		field := objectType.Field(i)

		// Skip gorm.Model fields to avoid issues with comparing timestamp fields
		if field.Name != "Model" {
			value1 := firstRecordValues.FieldByName(field.Name).Interface()
			value2 := secondRecordValues.FieldByName(field.Name).Interface()

			if value1 != value2 {
				equal = false
				unequalFields = append(unequalFields, field.Name)
			}
		}
	}

	return unequalFields, equal
}
