package app

import (
	"fmt"
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
	envconfig.MustProcess("API", &api)

	var db Database
	envconfig.MustProcess("DB", &db)

	fmt.Printf("%#v\n", db)

	return Cfg{api, db}
}
