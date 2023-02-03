package config

import (
	"log"

	"github.com/spf13/viper"
)

var ConfigVars *configVars

func InitEnvConfigs() {
	ConfigVars = loadEnvVariables()
}

// struct to map env values
type configVars struct {
	JWT_Key            string `mapstructure:"JWT_KEY"`
	DatabaseConnection string `mapstructure:"DATABASE_CONNECTION_STR"`
}

// Call to load the variables from env
func loadEnvVariables() (config *configVars) {

	viper.AddConfigPath(".")

	viper.SetConfigName(".env")

	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
