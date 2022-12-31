package app

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
)

func DB() *sql.DB {
	cfg := Config()

	dsn := fmt.Sprintf("postgres://%s:%d/%s?sslmode=%s&user=%s&password=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SslMode,
		cfg.Database.User,
		cfg.Database.Pass,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}
