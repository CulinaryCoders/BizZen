package database

import (
	"log"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: Add comment documentation (func Initialize)
func Initialize(connectionString string) *gorm.DB {

	dbInstance, dbError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")

	dbInstance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed!")

	return dbInstance
}
