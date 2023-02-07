package main

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	var dbConnectionString string = config.AppConfig.GetDBConnectionString()
	DB := database.Initialize(dbConnectionString)

	// Initialize router
	router := mux.NewRouter()

	// Initialize database handler
	dbHandler := handlers.NewHandler(DB)

	// Define routes
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// TODO: Create API table of contents or list of valid objects in API for root "/" route
		fmt.Fprint(writer, "Hello, World!")
	})

	router.HandleFunc("/register", dbHandler.CreateUser).Methods("POST")
	router.HandleFunc("/authenticate", dbHandler.Authenticate).Methods("POST")

	router.HandleFunc("/user/{id}", dbHandler.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", dbHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}}", dbHandler.DeleteUser).Methods("DELETE")

	log.Println("API is running!")
	log.Fatal(http.ListenAndServe(":8080", router))

}
