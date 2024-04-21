package main

import (
	"Savings/api"
	"Savings/pkg/datastore"
	"Savings/utils"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//https://gobyexample.com/command-line-subcommands

	migrateCmd := flag.NewFlagSet("migrate", flag.ExitOnError)

	routeCmd := flag.NewFlagSet("route", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'migrate' or 'route' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "migrate":
		migrateCmd.Parse(os.Args[2:])
		//fmt.Println("subcommand 'migrate'")
		//fmt.Println("  tail:", migrateCmd.Args())

		utils.SetupConfig(".")

		datastore.Init()
		defer datastore.Close()

		// Run the auto migration tool.
		if err := datastore.EntClient.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}

		fmt.Println("migrated")

	case "route":
		routeCmd.Parse(os.Args[2:])
		//fmt.Println("subcommand 'bar'")
		//fmt.Println("  tail:", routeCmd.Args())

		utils.SetupConfig(".")

		datastore.Init()
		defer datastore.Close()

		app := api.FiberServer()

		data, _ := json.MarshalIndent(app.GetRoutes(true), "", "  ")
		fmt.Print(string(data))

	default:
		fmt.Println("expected 'migrate' or 'route' subcommands")
		os.Exit(1)
	}
}
