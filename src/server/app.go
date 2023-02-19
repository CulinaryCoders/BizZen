package main

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/handlers"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// TODO:  Add documentation (type Application)
type Application struct {
	Router    *mux.Router
	DBHandler *handlers.DatabaseHandler
}

// TODO:  Add documentation (func Initialize)
func (app *Application) Initialize() {
	// Initialize main app database
	var dbConnectionString string = config.AppConfig.GetPostgresDBConnectionString()
	appDB := database.InitializePostgresDB(dbConnectionString)

	// Initialize cache db
	var cacheDSN string = config.AppConfig.GetRedisDBNetworkAddress()
	cacheDB := database.InitializeRedisDB(cacheDSN)

	// Initialize cookie store
	cookieStore := sessions.NewCookieStore([]byte("super-secret"))
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.Secure = true

	// Initialize DataHandler
	app.DBHandler = handlers.NewDatabaseHandler(appDB, cacheDB)

	// Initialize router and routes
	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

// TODO:  Add documentation (func initializeRoutes)
func (app *Application) initializeRoutes() {
	// Define routes
	app.Router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// TODO: Create API table of contents or list of valid objects in API for root "/" route
		fmt.Fprint(writer, "Hello, World!")
	})

	// User routes
	app.Router.HandleFunc("/register", app.DBHandler.CreateUser).Methods("POST")
	app.Router.HandleFunc("/authenticate", app.DBHandler.Authenticate).Methods("POST")
	app.Router.HandleFunc("/user/{id}", app.DBHandler.GetUser).Methods("GET")
	app.Router.HandleFunc("/user/{id}", app.DBHandler.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/user/{id}}", app.DBHandler.DeleteUser).Methods("DELETE")

	// Business routes
	app.Router.HandleFunc("/business", app.DBHandler.CreateBusiness).Methods("POST")

}

// TODO:  Add documentation (func Run)
func (app *Application) Run(networkAddress string) {
	log.Printf("Server is now listening on: %s", networkAddress)
	log.Fatal(http.ListenAndServe(networkAddress, app.Router))
}
