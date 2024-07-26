package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig(fileName string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(fileName)
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")
	err := config.ReadInConfig()

	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	return config
}
