package main

import (
	"server/config"
	"testing"
)

// Perform basic environment setup for testing
func TestMain(m *testing.M) {
	var testDBName string = config.AppConfig.APP_TEST_DB_NAME

	app := Application{}
	app.Initialize(testDBName)
}
