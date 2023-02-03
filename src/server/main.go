package main

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/handlers"
	"server/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Database
	config.InitEnvConfigs()
	DB := database.Init(config.ConfigVars.DatabaseConnection)
	router := mux.NewRouter()
	db := handlers.NewHandler(DB)

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, World!")
	})

	router.HandleFunc("/register", db.RegisterUser).Methods("POST")
	router.HandleFunc("/login", db.Authenticate).Methods("POST")

	router.HandleFunc("/customer", middlewares.Authorize(db.Customer)).Methods("GET")

	router.HandleFunc("/finduser/{email}", db.FindUser).Methods("GET")
	router.HandleFunc("/updateuser/{email}", db.UpdateUser).Methods("POST")
	router.HandleFunc("/deleteuser/{email}", db.DeleteUser).Methods("DELETE")

	log.Println("API is running!")
	log.Fatal(http.ListenAndServe(":8080", router))

}
