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

/*
*Description*

func InitializePostgresDB

Initializes and connects to a Postgres database using the provided command.

*Parameters*

	connectionString  <string>

		The connection command used to establish the connection with the Postgres database.

	debug  <bool>

		A flag variable used to trigger additional logging in the &gorm.Config definition for the returned *gorm.DB instance.

*Returns*

	_  <*gorm.DB>

		The Postgres database instance.
*/
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

/*
*Description*

func InitializeRedisDB

Initializes and connects to a Redis database using the provided dsn.

*Parameters*

	dsn  <string>

		The host address and port number the Redis DB is hosted on in 'host:port' format.

*Returns*

	_  <*redis.Client>

		The initialized redis client.
*/
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

/*
*Description*

func FormatAllTables

Performs a complete refresh of the specified database. All tables are dropped/deleted from the database and then new tables are created for all DB models.

*Parameters*

	db <*gorm.DB>

		The database instance that will be formatted and refreshed.

*Returns*

	N/A (None)
*/
func FormatAllTables(db *gorm.DB) {
	dropAllTables(db)
	setupTables(db)
}

/*
*Description*

func setupTables

Executes gorm AutoMigrate function for all DB models.

*Parameters*

	db <*gorm.DB>

		The database instance where the tables will be created/updated.

*Returns*

	N/A (None)
*/
func setupTables(db *gorm.DB) {
	db.AutoMigrate(&models.User{},
		// &models.Address{},
		// &models.ContactInfo{},
		&models.Business{},
		// &models.Office{},
		&models.Service{},
		// &models.ServiceOffering{},
		&models.Appointment{})
}

/*
*Description*

func dropAllTables

Drops all of the tables present in the specified database instance.

*Parameters*

	db  <*gorm.DB>

		The database instance that will have all tables dropped from it.

*Returns*

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
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

/*
*Description*

func getListOfDBTables

Returns the list of names of all the tables present in the specified database instance.

*Parameters*

	db  <*gorm.DB>

		The database instance that will have all tables dropped from it.

*Returns*

	tableNames  <[]string>

		The list of table names present in the specified database instance.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func getListOfDBTables(db *gorm.DB) (tableNames []string, err error) {
	if err := db.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tableNames).Error; err != nil {
		return tableNames, err
	} else {
		return tableNames, nil
	}
}
