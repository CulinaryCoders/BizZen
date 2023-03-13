package tests

import (
	"os"
	"server/config"
	"server/database"
	"testing"

	"gorm.io/gorm"
)

var testAppDB *gorm.DB

func TestMain(m *testing.M) {
	var testDBName string = config.AppConfig.APP_TEST_DB_NAME
	var dbConnectionString string = config.AppConfig.GetPostgresDBConnectionString(testDBName)
	var debug bool = config.AppConfig.DEBUG_MODE
	testAppDB = database.InitializePostgresDB(dbConnectionString, debug)

	os.Exit(m.Run())
}
