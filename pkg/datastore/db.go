package datastore

import (
	"Savings/ent"
	"context"
	"github.com/spf13/viper"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var EntClient *ent.Client

func Init() {
	dsn := viper.GetString("DB_DSN")
	env := strings.ToUpper(viper.GetString("APP_ENV"))
	if env != "TEST" && dsn == "" {
		panic("database connection not set")
	}

	if env == "TEST" {
		EntClient, err := ent.Open("sqlite3", "file:storage/data.db?cache=shared&_fk=1")
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
		defer EntClient.Close()

		// Run the auto migration tool.
		if err := EntClient.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		return
	}

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer EntClient.Close()

	EntClient = client
}
