package database

import (
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"log"
)

func ApplyMigrations() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:yourpassword@localhost/mydatabase?sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations:%v", err)
	}

	log.Println("Migrations applied successfully")
}

func RollbackMigrations() {
	m, err := migrate.New(
		"file://migrations",

		"postgres://postgres:yourpassword@localhost/mydatabase?sslmode=disable",
	)

	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to rollback migrations: %v", err)
	}

	log.Println("Migrations rolled back successfully")
}
