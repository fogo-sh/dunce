package cmd

import (
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/database"
)

var dbVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Fetch the current database version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := database.NewMigrateInstance(config.DBPath)
		checkError(err, "Error setting up migrations")

		version, dirty, err := m.Version()
		checkError(err, "Error fetching database version")

		slog.Info(fmt.Sprintf("Version %d", version), "dirty", dirty)
	},
}

func init() {
	dbCmd.AddCommand(dbVersionCmd)
}
