package main

import (
	"log"
	"server/config"
	"server/database"
	"server/handlers"
	"server/sample_data"
)

func main() {
	var prodDBName string = config.AppConfig.APP_DB_NAME

	// Initialize application
	handlers.App.Initialize(prodDBName)

	// Format and table creation refresh of database
	database.FormatAllTables(handlers.App.AppDB)

	// Load sample data from JSON files
	err := sample_data.LoadJSONSampleData(handlers.App.AppDB)
	if err != nil {
		log.Fatal(err)
	}

	// Launch application instance
	handlers.App.Run(config.AppConfig.GetAPIServerNetworkAddress())

	// TODO:  Determine what OS application is running on and load sample records into DB using REST calls
	// if runtime.GOOS == "windows" {

	// Trigger Powershell script
	// 	log.Printf("Encountered Windows OS:  %s", runtime.GOOS)

	// 	serverExe, err := os.Executable()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	serverExePath := filepath.Dir(serverExe)

	// 	ps1Path := filepath.Join(serverExePath, "sample_data", "load-sample-data.ps1")
	// 	ps1Arg := fmt.Sprintf("-Command %s", ps1Path)
	// 	cmd := exec.Command("Powershell", "-NoLogo", "-NoExit", "-NoProfile", ps1Arg)

	// 	fmt.Println(cmd.String())
	// 	var userExit string
	// 	fmt.Scanln(userExit)
	// 	if err := cmd.Run(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// Trigger shell script
	// 	log.Printf("Encountered Unix OS:  %s", runtime.GOOS)
	// }
}
