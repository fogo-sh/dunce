package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/database"
)

var dbDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Completely empty the current database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(config.DBPath)
		checkError(err, "Error setting up migrations")

		slog.Info("Dropping database...")
		err = m.Drop()
		checkError(err, "Error dropping database")

		slog.Info("Database dropped!")
	},
}

func init() {
	dbCmd.AddCommand(dbDropCmd)
}
