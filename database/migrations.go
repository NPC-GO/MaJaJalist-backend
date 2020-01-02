package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func migration() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@database:5432/MaJaJalist?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		err := m.Down()
		if err != nil {
			log.Fatal(err)
		}
		err = m.Up()
		if err != nil {
			log.Fatal(err)
		}
	}
}
