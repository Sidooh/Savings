package main

import (
	"Savings/api"
	"Savings/pkg/datastore"
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
	datastore.Init()
	defer datastore.Close()

	//for i := 0; i < 100; i++ {
	//	datastore.EntClient.PersonalAccount.Create().
	//		SetAccountID(uint64(utils.RandomInt(1, 1000))).
	//		SetType(utils.RandomString(6)).
	//		SetBalance(float32(utils.RandomInt(0, 100))).
	//		Save(context.Background())
	//}

	//cache.Init()
	//clients.Init()

	api.Server()
}
