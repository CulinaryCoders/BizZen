package main

import (
	"server/config"
	"server/database"
	"server/handlers"
)

func main() {
	var prodDBName string = config.AppConfig.APP_DB_NAME

	handlers.App.Initialize(prodDBName)
	database.FormatAllTables(handlers.App.AppDB)
	handlers.App.Run(config.AppConfig.GetAPIServerNetworkAddress())
}
