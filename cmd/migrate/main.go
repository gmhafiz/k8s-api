package main

import (
	"flag"
	"log"

	app "github.com/gmhafiz/k8s-api"
)

func main() {
	log.Println("starting migrate...")

	var migrateCommand string
	flag.StringVar(&migrateCommand, "go", "up", "-go up")
	flag.Parse()

	cfg := app.Config()

	db := app.DB(cfg.Database)

	migrator := app.Migrator(db)

	if migrateCommand == "up" {
		migrator.Up()
	} else if migrateCommand == "down" {
		migrator.Down()
	} else {
		log.Println("operation not supported")
	}
}
