package tests

import (
	"fmt"
	"server/config"
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestInitializePostgresDB

Tests the InitializePostgresDB method to confirm that it successfully establishes a connection with the specified database.
*/
func TestInitializePostgresDB(t *testing.T) {
	var unitTestDBName string = config.AppConfig.APP_TEST_DB_NAME
	var dbConnectionString string = config.AppConfig.GetPostgresDBConnectionString(unitTestDBName)

	unitTestGormDB := models.InitializePostgresDB(dbConnectionString, config.Debug)

	pingTestLogMessage := fmt.Sprintf("Attempting to ping Postgres database (%s) to confirm successful connection.", unitTestDBName)
	t.Run(pingTestLogMessage, func(t *testing.T) {
		//  Get DB instance from GORM object
		unitTestDB, err := unitTestGormDB.DB()
		if err != nil {
			t.Errorf("Unable to initialize DB instance from GORM object (database name:  %s)  --  %s", unitTestDBName, err)
		}
		//  Attempt to ping database to confirm that connection was successfully established
		err = unitTestDB.Ping()
		if err != nil {
			t.Errorf("Unable to ping Postgres database (%s)  --  %s", unitTestDBName, err)
		}
	})
}

/*
*Description*

func TestFormatAllTables

Tests the FormatAllTables method to confirm that all object tables exist in the database and have 0 rows present after DB is formatted
*/
func TestFormatAllTables(t *testing.T) {
	expectedRowCount := 0
	tableList := []string{
		"users",
		"businesses",
		"services",
		"appointments",
		"invoices",
	}

	models.FormatAllTables(testAppDB)

	for _, tableName := range tableList {

		t.Run(fmt.Sprintf("Confirming %s has %d rows after DB format.", tableName, expectedRowCount), func(t *testing.T) {
			var actualRowCount int64
			testAppDB.Raw("SELECT * from ?", tableName).Count(&actualRowCount)

			assert.Equal(t, int(actualRowCount), expectedRowCount, "Actual row count (%d) matches expected row count (%d)", actualRowCount, expectedRowCount)
		})

	}

}
