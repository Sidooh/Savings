package utils

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Environment string
	JwtKey      string
	Port        int
}

func SetupConfig(path string) {
	// Set the path to look for the configurations file
	viper.AddConfigPath(path)

	// Set the file name of the configurations file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			// ignore Config file not found; report all others
			log.Fatal("Fatal error: ", err)
		}
	}
}
