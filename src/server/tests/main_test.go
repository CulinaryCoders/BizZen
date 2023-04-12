package tests

import (
	"os"
	"server/config"
	"server/models"
	"testing"

	"gorm.io/gorm"
)

var testAppDB *gorm.DB

func TestMain(m *testing.M) {
	var testDBName string = config.AppConfig.APP_TEST_DB_NAME
	var dbConnectionString string = config.AppConfig.GetPostgresDBConnectionString(testDBName)
	testAppDB = models.InitializePostgresDB(dbConnectionString, config.Debug)

	os.Exit(m.Run())
}
