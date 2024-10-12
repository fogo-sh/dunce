package cmd

import (
	"errors"
	"log/slog"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/database"
)

var dbStepCmd = &cobra.Command{
	Use: "step [number of steps]",
	Example: `  Downgrade by one version: netenvelope db step -- -1
  Upgrade by 3 versions: netenvelope db step 3`,
	Short: "Perform a relative migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(config.DBPath)
		checkError(err, "Error setting up migrations")

		stepCountStr := args[0]

		stepCount, err := strconv.Atoi(stepCountStr)
		checkError(err, "Error parsing step count")

		slog.Info("Performing database migration...")
		err = m.Steps(stepCount)
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				slog.Info("Database is already at desired version.")
				return
			} else {
				slog.Error("Error upgrading database", err)
			}
		}

		slog.Info("Database migrations complete!")
	},
}

func init() {
	dbCmd.AddCommand(dbStepCmd)
}
