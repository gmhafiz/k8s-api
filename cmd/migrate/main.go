package main

import (
	app "github.com/gmhafiz/k8s-api"
	"log"
)

func main() {
	log.Println("starting migrate...")
	db := app.DB()

	migrator := app.Migrator(db)
	migrator.Up()
}
