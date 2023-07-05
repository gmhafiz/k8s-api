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

	_, err = db.Exec("SELECT true")
	if err != nil {
		log.Panic(err)
	}

	return db
}
