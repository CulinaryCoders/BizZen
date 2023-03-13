package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Read in environmental variables used throughout application
var AppConfig *Configuration = InitializeConfig()

// struct to map env values
type Configuration struct {
	JWT_SIGNING_KEY   []byte `mapstructure:"JWT_SIGNING_KEY"`
	APP_DB_NAME       string `mapstructure:"APP_DB_NAME"`
	APP_TEST_DB_NAME  string `mapstructure:"APP_TEST_DB_NAME"`
	APP_DB_USER       string `mapstructure:"APP_DB_USER"`
	APP_DB_PASSWORD   string `mapstructure:"APP_DB_PASSWORD"`
	APP_DB_HOST       string `mapstructure:"APP_DB_HOST"`
	APP_DB_PORT       int    `mapstructure:"APP_DB_PORT"`
	APP_CACHE_DB_HOST string `mapstructure:"APP_CACHE_DB_HOST"`
	APP_CACHE_DB_PORT int    `mapstructure:"APP_CACHE_DB_PORT"`
	API_SERVER_HOST   string `mapstructure:"API_SERVER_HOST"`
	API_SERVER_PORT   int    `mapstructure:"API_SERVER_PORT"`
	FRONTEND_HOST     string `mapstructure:"FRONTEND_HOST"`
	FRONTEND_PORT     int    `mapstructure:"FRONTEND_PORT"`
	DEBUG_MODE        bool   `mapstructure:"DEBUG_MODE"`
}

// Initialize method creates and initializes new Configuration object
func InitializeConfig() (config *Configuration) {
	return loadEnvironmentVariables()
}

// loadEnvironmentVariables reads in the variables found in the 'server/config/config.json' file and returns a struct with those variables loaded as properties
func loadEnvironmentVariables() (config *Configuration) {
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.SetConfigName("config.json")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}

// GetSigningKey returns the JWT_SIGNING_KEY environmental variable defined in the calling Configuration object
func (config *Configuration) GetSigningKey() []byte {
	return config.JWT_SIGNING_KEY
}

// GetPostgresDBConnectionString returns the formatted connection string for Postgres database connections
func (config *Configuration) GetPostgresDBConnectionString(appDBName string) string {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		config.APP_DB_HOST,
		config.APP_DB_USER,
		config.APP_DB_PASSWORD,
		appDBName,
		config.APP_DB_PORT)

	return connectionString
}

// GetAPIServerNetworkAddress returns the full network address ("host:port") that will serve the API
func (config *Configuration) GetAPIServerNetworkAddress() string {
	return getNetworkAddress(
		config.API_SERVER_HOST,
		config.API_SERVER_PORT)
}

// GetRedisDBNetworkAddress returns DSN for Redis DB in "host:port" string format
func (config *Configuration) GetRedisDBNetworkAddress() string {
	return getNetworkAddress(
		config.APP_CACHE_DB_HOST,
		config.APP_CACHE_DB_PORT)
}

// GetFrontendNetworkAddress returns the full network address ("host:port") that will serve the frontend Angular application
func (config *Configuration) GetFrontendNetworkAddress() string {
	return getNetworkAddress(
		config.FRONTEND_HOST,
		config.FRONTEND_PORT)
}

// getNetworkAddress takes in host address and port number and returns the network address (a.k.a. DSN) in "host:port" string format
func getNetworkAddress(host string, port int) string {
	var networkAddress string = fmt.Sprintf("%s:%d",
		host,
		port)

	return networkAddress
}
