package database

import (
	"fmt"
	"log"
	"server/models"

	"github.com/go-redis/redis/v7"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO: Add comment documentation (func InitializePostgresDB)
func InitializePostgresDB(connectionString string) *gorm.DB {

	dbInstance, dbError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")

	dbInstance.AutoMigrate(&models.User{},
		&models.Address{},
		&models.ContactInfo{},
		&models.Business{},
		&models.Office{},
		&models.Service{},
		&models.ServiceOffering{},
		&models.Appointment{})

	log.Println("Database Migration Completed!")

	return dbInstance
}

// TODO: Add comment documentation (func InitializeRedisDB)
func InitializeRedisDB(dsn string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}

	return client
}
