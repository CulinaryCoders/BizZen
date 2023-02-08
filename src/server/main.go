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
	// Initialize Database
	config.InitEnvConfigs()
	DB := database.Init(config.ConfigVars.DatabaseConnection)
	redis := handlers.Init()
	router := mux.NewRouter()
	db := handlers.NewHandler(DB, redis)
	//controllers.Init()

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello, World!")
	})

	router.HandleFunc("/register", db.RegisterUser).Methods("POST")
	router.HandleFunc("/newlogin", db.Authenticate).Methods("POST")
	router.HandleFunc("/login", db.OldLogin).Methods("POST") //login with previous authentication method

	router.HandleFunc("/customer", db.Authorize(db.Customer)).Methods("GET")
	router.HandleFunc("/business", db.Authorize(db.Business)).Methods("GET")
	//router.HandleFunc("/getuser", db.GetUserDetails).Methods("GET")
	router.HandleFunc("/finduser/{email}", db.Authorize(db.FindUser)).Methods("GET")
	router.HandleFunc("/updateuser/{email}", db.UpdateUser).Methods("POST")
	router.HandleFunc("/deleteuser/{email}", db.DeleteUser).Methods("DELETE")

	log.Println("API is running!")
	log.Fatal(http.ListenAndServe(":8080", router))

}
