package main

import (
	"log"
	"server/config"
	"server/handlers"
	"server/models"
	"server/sample_data"
)

func main() {
	var prodDBName string = config.AppConfig.APP_DB_NAME

	// Initialize application
	handlers.App.Initialize(prodDBName)

	// Format and table creation refresh of database
	models.FormatAllTables(handlers.App.AppDB)

	// Load sample data from JSON files
	err := sample_data.LoadJSONSampleData(handlers.App.AppDB)
	if err != nil {
		log.Fatal(err)
	}

	// Launch application instance
	handlers.App.Run(config.AppConfig.GetAPIServerNetworkAddress())
}
