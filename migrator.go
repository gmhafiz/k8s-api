package app

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Migrate struct {
	DB *sql.DB
}

func Migrator(db *sql.DB) *Migrate {
	m := &Migrate{
		DB: db,
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Panic(err)
	}
	goose.SetBaseFS(embedMigrations)

	return m
}

func (m *Migrate) Up() {
	if err := goose.Up(m.DB, "migrations"); err != nil {
		log.Panic(err)
	}

	if err := goose.Version(m.DB, "migrations"); err != nil {
		log.Panic(err)
	}
}

func (m *Migrate) Down() {
	if err := goose.Down(m.DB, "migrations"); err != nil {
		log.Panic(err)
	}

	if err := goose.Version(m.DB, "migrations"); err != nil {
		log.Panic(err)
	}
}
