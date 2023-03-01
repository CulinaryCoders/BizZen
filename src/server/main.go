package main

import (
	"server/config"
	"server/handlers"
)

func main() {
	var prodDBName string = config.AppConfig.APP_DB_NAME

	handlers.App.Initialize(prodDBName)
	handlers.App.Run(config.AppConfig.GetAPIServerNetworkAddress())
}
