package main

import (
	"log"
	"server/config"
)

func main() {
	app := Application{}
	app.Initialize()
	app.Run(config.AppConfig.GetAPIServerNetworkAddress())

	log.Println("API is running!")
}
