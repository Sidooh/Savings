package main

import (
	"Savings/api"
	"Savings/pkg/datastore"
	"Savings/utils"
	"Savings/utils/logger"
	"context"
	"fmt"
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
	//cache.Init()
	//clients.Init()

	api.Server()

	fmt.Println(datastore.EntClient.PersonalAccount.Query().All(context.Background()))
}
