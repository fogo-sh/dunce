package database

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"
)

//go:embed migrations
var migrationsFs embed.FS

func NewMigrateInstance(connectionString string) (*migrate.Migrate, error) {
	db, err := sql.Open("sqlite", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return nil, fmt.Errorf("error creating migrate driver: %w", err)
	}

	migrationSource, err := iofs.New(migrationsFs, "migrations")
	if err != nil {
		return nil, fmt.Errorf("error creating migration source: %w", err)
	}

	migrateInstance, err := migrate.NewWithInstance(
		"iofs", migrationSource,
		"sqlite", driver,
	)
	if err != nil {
		return nil, err
	}

	return migrateInstance, nil
}
