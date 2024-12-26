package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/bot"
)

var shouldMigrate bool

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Discord bot",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if shouldMigrate {
			err := performDatabaseUpgrade()
			if err != nil {
				slog.Error("Error upgrading database", "error", err)
				return
			}
		}

		err := bot.Run(config)
		if err != nil {
			slog.Error("Error running Dunce", "error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVar(&shouldMigrate, "migrate", false, "Run database migrations before starting the bot")
}
