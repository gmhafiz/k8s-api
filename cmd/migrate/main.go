package main

import (
	"log"

	app "github.com/gmhafiz/k8s-api"
)

func main() {
	log.Println("starting migrate...")

	cfg := app.Config()
	db := app.DB(cfg.Database)

	migrator := app.Migrator(db)
	migrator.Up()
}
