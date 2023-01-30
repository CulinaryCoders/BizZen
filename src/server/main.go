package main

import (
	"fmt"
	"log"
	"net/http"
	"server/database"
	"server/handlers"
	"server/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Database
	DB := database.Init()
	router := mux.NewRouter()
	h := handlers.New(DB)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	router.HandleFunc("/register", h.RegisterUser).Methods("POST")
	router.HandleFunc("/login", h.LogIn).Methods("POST")

	router.HandleFunc("/customer", middlewares.Auth(h.Customer)).Methods("GET")

	router.HandleFunc("/finduser/{email}", h.FindUser).Methods("GET")
	router.HandleFunc("/updateuser/{email}", h.UpdateUser).Methods("POST")
	router.HandleFunc("/deleteuser/{email}", h.DeleteUser).Methods("DELETE")

	log.Println("API is running!")
	http.ListenAndServe(":8080", router)

}
