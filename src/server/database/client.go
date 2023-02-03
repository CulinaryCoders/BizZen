package database

import (
	"log"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

// TODO: Add comment documentation (func Init)
func Init(connectionString string) *gorm.DB {

	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")

	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")

	return Instance
}
