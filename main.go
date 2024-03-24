package main

import (
	"Savings/utils"
	"Savings/utils/logger"
	"github.com/spf13/viper"
)

func main() {
	utils.SetupConfig(".")

	jwtKey := viper.GetString("JWT_KEY")
	if len(jwtKey) == 0 {
		panic("JWT_KEY is not set")
	}

	logger.Init()
	//datastore.Init()
	//cache.Init()
	//clients.Init()

	logger.Log.Info("started")

}
