package database

import (
	"errors"
	"fmt"
	"log"
	"server/models"

	"github.com/go-redis/redis/v7"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TODO: Add comment documentation (func InitializePostgresDB)
func InitializePostgresDB(connectionString string, debug bool) *gorm.DB {
	var gormConfig *gorm.Config

	if debug {
		gormConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	} else {
		gormConfig = &gorm.Config{}
	}

	dbInstance, dbError := gorm.Open(postgres.Open(connectionString), gormConfig)

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")

	setupTables(dbInstance)

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

// TODO: Add comment documentation (func FormatAllTables)
func FormatAllTables(db *gorm.DB) {
	dropAllTables(db)
	setupTables(db)
}

// TODO: Add comment documentation (func setupTables)
func setupTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{},
		&models.Address{},
		&models.ContactInfo{},
		&models.Business{},
		&models.Office{},
		&models.Service{},
		&models.ServiceOffering{},
		&models.Appointment{})
}

// TODO: Add comment documentation (func dropAllTables)
func dropAllTables(db *gorm.DB) error {
	tableNames, err := getListOfDBTables(db)
	if err != nil {
		return err
	}

	var hasDropTableErrors bool = false
	errorTables := make(map[string]error)
	for _, tableName := range tableNames {
		err := db.Migrator().DropTable(tableName)
		if err != nil {
			hasDropTableErrors = true
			errorTables[tableName] = err
		}
	}

	if hasDropTableErrors {
		var customErrorMessage string = fmt.Sprintf("Unable to drop the following table(s) (func DropAllTables):\n\n%v", errorTables)
		return errors.New(customErrorMessage)
	} else {
		return nil
	}
}

// TODO: Add comment documentation (func getListOfDBTables)
func getListOfDBTables(db *gorm.DB) (tableNames []string, err error) {
	if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tableNames).Error; err != nil {
		return tableNames, err
	} else {
		return tableNames, nil
	}
}
