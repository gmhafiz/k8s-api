package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func DB(cfg Database) *sql.DB {
	dsn := fmt.Sprintf("postgres://%s:%d/%s?sslmode=%s&user=%s&password=%s",
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SslMode,
		cfg.User,
		cfg.Pass,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil
	}

	//db := stdlib.OpenDB(pgx.ConnConfig{
	//	Config: pgconn.Config{
	//		Host:      cfg.Host,
	//		Port:      cfg.Port,
	//		Database:  cfg.Name,
	//		User:      cfg.User,
	//		Password:  cfg.Pass,
	//		TLSConfig: nil,
	//		//ConnectTimeout:  0,
	//		//DialFunc:        nil,
	//		//LookupFunc:      nil,
	//		//BuildFrontend:   nil,
	//		//RuntimeParams:   nil,
	//		//KerberosSrvName: "",
	//		//KerberosSpn:     "",
	//		//Fallbacks:       nil,
	//		//ValidateConnect: nil,
	//		//AfterConnect:    nil,
	//		//OnNotice:        nil,
	//		//OnNotification:  nil,
	//	},
	//	//Tracer:                   nil,
	//	//StatementCacheCapacity:   0,
	//	//DescriptionCacheCapacity: 0,
	//	//DefaultQueryExecMode:     0,
	//})

	_, err = db.Exec("SELECT true")
	if err != nil {
		log.Panic(err)
	}

	return db
}
