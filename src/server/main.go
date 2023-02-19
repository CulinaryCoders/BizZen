package main

import (
	"server/config"
)

func main() {
	app := Application{}
	app.Initialize()
	app.Run(config.AppConfig.GetAPIServerNetworkAddress())
}
