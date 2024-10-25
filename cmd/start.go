package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/bot"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Discord bot",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := bot.Run(config)
		if err != nil {
			slog.Error("Error running Dunce", "error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
