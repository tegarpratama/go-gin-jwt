package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST       string
	DB_USER       string
	DB_PASSWORD   string
	DB_DATABASE   string
	DB_PORT       string
	PORT          string
	CLIENT_ORIGIN string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&ENV)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Configuration loaded")
}
