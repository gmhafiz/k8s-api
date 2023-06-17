package app

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Cfg struct {
	Api
	Database
}
type Api struct {
	Name string `default:"api"`
	Host string `default:"0.0.0.0"`
	Port uint16 `default:"3080"`
}

type Database struct {
	Host    string
	Port    uint16
	Name    string
	User    string
	Pass    string
	SslMode string `default:"disable"`
}

func Config() Cfg {
	log.Println("reading env")

	var api Api
	err := envconfig.Process("API", &api)
	if err != nil {
		log.Println("no API config is found")
	}

	var db Database
	err = envconfig.Process("DB", &db)
	if err != nil {
		log.Println("no DB config is found")
	}

	return Cfg{api, db}
}
