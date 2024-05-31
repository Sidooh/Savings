package main

import (
	"Savings/api"
	"Savings/pkg/clients"
	"Savings/pkg/datastore"
	"Savings/utils"
	"Savings/utils/cache"
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
	datastore.Init()
	defer datastore.Close()

	cache.Init()
	clients.Init()

	api.Server()
}
