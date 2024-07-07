package config

import (
	"log"

	"github.com/G-MAKROGLOU/ais/pkg/utils"
	"github.com/spf13/viper"
)

// AppConfig holds the values of all the configuration required by the application
var AppConfig *Config

// LoadEnv loads the environment variables for each environment.
// For development it will load the values from .env
// For production or staging it will load the values from the system
// If the value is missing, it will try to load the values from the system
func LoadEnv() {
	viper.AutomaticEnv()

	env := viper.GetString("AIS_ENV")

	if env == "development" {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln("Failed to load development environment variables", err)
		}
	}

	AppConfig = &Config{
		Port:             viper.GetString("PORT"),
		Env:              viper.GetString("AIS_ENV"),
		PostgresHost:     viper.GetString("POSTGRES_HOST"),
		PostgresUser:     viper.GetString("POSTGRES_USER"),
		PostgresPassword: viper.GetString("POSTGRES_PASSWORD"),
		PostgresPort:     viper.GetString("POSTGRES_PORT"),
		PostgresDB:       viper.GetString("POSTGRES_DB"),
		RabbitMQHost:     viper.GetString("RABBITMQ_HOST"),
		RabbitMQUser:     viper.GetString("RABBITMQ_USER"),
		RabbitMQPassword: viper.GetString("RABBITMQ_PASSWORD"),
		RabbitMQPort:     viper.GetString("RABBITMQ_PORT"),
		RabbitMQTopic:    viper.GetString("RABBITMQ_TOPIC"),
	}
	utils.PrettyPrint(AppConfig)

}
