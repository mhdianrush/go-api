package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	PORT        string
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	logger := logrus.New()

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
	err = viper.Unmarshal(&ENV)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Load Server Successfully")
}
