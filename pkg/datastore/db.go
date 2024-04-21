package datastore

import (
	"Savings/ent"
	"context"
	"fmt"
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
		if dsn == "" {
			dsn = "file:storage/data.db?cache=shared&_fk=1"
		}
		client, err := ent.Open("sqlite3", dsn)
		if err != nil {
			panic(fmt.Sprintf("failed opening connection to db: %v", err))
		}

		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		EntClient = client

		return
	}

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to db: %v", err))
	}

	// Ensure DB is live
	_, err = client.ExecContext(context.Background(), "SELECT 1;")
	if err != nil {
		panic(err)
	}

	//// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	EntClient = client
}

func Close() {
	EntClient.Close()
}
