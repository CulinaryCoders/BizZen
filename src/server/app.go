package main

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// TODO:  Add documentation (type Application)
type Application struct {
	Router    *mux.Router
	DB        *gorm.DB
	DBHandler *handlers.Handler
}

// TODO:  Add documentation (func Initialize)
func (app *Application) Initialize() {
	// Initialize database
	var dbConnectionString string = config.AppConfig.GetDBConnectionString()
	app.DB = database.Initialize(dbConnectionString)

	// Initialize database handler
	app.DBHandler = handlers.NewHandler(app.DB)

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

	app.Router.HandleFunc("/register", app.DBHandler.CreateUser).Methods("POST")
	app.Router.HandleFunc("/authenticate", app.DBHandler.Authenticate).Methods("POST")

	app.Router.HandleFunc("/user/{id}", app.DBHandler.GetUser).Methods("GET")
	app.Router.HandleFunc("/user/{id}", app.DBHandler.UpdateUser).Methods("PUT")
	app.Router.HandleFunc("/user/{id}}", app.DBHandler.DeleteUser).Methods("DELETE")
}

// TODO:  Add documentation (func Run)
func (app *Application) Run(networkAddress string) {
	log.Fatal(http.ListenAndServe(networkAddress, app.Router))
}
