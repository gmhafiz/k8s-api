package app

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed database/migrations/*.sql
var embedMigrations embed.FS

type Migrate struct {
	DB *sql.DB
}

func Migrator(db *sql.DB) *Migrate {

	m := &Migrate{
		DB: db,
	}
	goose.SetBaseFS(embedMigrations)

	return m
}

func (m *Migrate) Up() {
	if err := goose.Up(m.DB, "database/migrations"); err != nil {
		log.Panic(err)
	}
}
