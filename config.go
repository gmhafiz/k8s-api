package app

import (
	"log"
	"os"
	"strconv"
)

type Cfg struct {
	Api
	Database
}
type Api struct {
	Name string
	Host string
	Port int
}

type Database struct {
	Host    string
	Port    int
	Name    string
	User    string
	Pass    string
	SslMode string
}

func Config() Cfg {
	log.Println("reading env")

	apiName := os.Getenv("API_NAME")
	apiHost := os.Getenv("API_HOST")
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		apiPort = 3080
	}
	api := Api{
		Name: apiName,
		Host: apiHost,
		Port: apiPort,
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5432
	}
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	db := Database{
		Host:    dbHost,
		Port:    dbPort,
		Name:    dbName,
		User:    dbUser,
		Pass:    dbPass,
		SslMode: dbSSLMode,
	}

	return Cfg{api, db}
}
