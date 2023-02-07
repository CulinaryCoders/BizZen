package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// struct to map env values
type Configuration struct {
	JWT_SIGNING_KEY []byte `mapstructure:"JWT_SIGNING_KEY"`
	APP_DB_NAME     string `mapstructure:"APP_DB_NAME"`
	APP_DB_USER     string `mapstructure:"APP_DB_USER"`
	APP_DB_PASSWORD string `mapstructure:"APP_DB_PASSWORD"`
	APP_DB_HOST     string `mapstructure:"APP_DB_HOST"`
	APP_DB_PORT     int    `mapstructure:"APP_DB_PORT"`
}

// Initialize method creates and initializes new Configuration object
func Initialize() (config *Configuration) {
	return loadEnvironmentVariables()
}

// loadEnvironmentVariables reads in the variables found in the 'server/config/config.json' file and returns a struct with those variables loaded as properties
func loadEnvironmentVariables() (config *Configuration) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}

// GetSigningKey returns the JWT_SIGNING_KEY environmental variable defined in the calling Configuration object
func (config *Configuration) GetSigningKey() []byte {
	return config.JWT_SIGNING_KEY
}

// GetDBConnectionString returns the formatted connection string for Postgres database connections
func (config *Configuration) GetDBConnectionString() string {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.APP_DB_HOST,
		config.APP_DB_USER,
		config.APP_DB_PASSWORD,
		config.APP_DB_NAME,
		config.APP_DB_PORT)

	return connectionString
}
