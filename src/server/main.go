package main

import (
	"server/config"
)

func main() {
	var prodDBName string = config.AppConfig.APP_DB_NAME

	app := Application{}
	app.Initialize(prodDBName)
	app.Run(config.AppConfig.GetAPIServerNetworkAddress())
}
