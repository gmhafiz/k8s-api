package app

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
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
	Driver  string `default:"postgres"`
	Host    string `default:"localhost"`
	Port    uint16 `default:"5432"`
	Name    string `default:"postgres"`
	User    string `default:"postgres"`
	Pass    string `default:"password"`
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
