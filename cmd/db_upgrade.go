package cmd

import (
	"errors"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/database"
)

func performDatabaseUpgrade() error {
	m, err := database.NewMigrateInstance(config.DBPath)
	if err != nil {
		return err
	}

	slog.Info("Performing database upgrades...")
	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("Database is already up to date!")
			return nil
		} else {
			return err
		}
	}

	slog.Info("Database upgrades complete!")
	return nil
}

var dbUpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"up"},
	Short:   "Perform all database upgrades",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := performDatabaseUpgrade()
		checkError(err, "Error upgrading database")
	},
}

func init() {
	dbCmd.AddCommand(dbUpgradeCmd)
}
