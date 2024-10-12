package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"

	"github.com/fogo-sh/dunce/bot"
)

var config bot.Config

var rootCmd = &cobra.Command{
	Use:   "dunce",
	Short: "Dunce",
}

func Execute() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Failed to load .env file", "error", err)
		os.Exit(1)
	}

	err = envconfig.Process("dunce", &config)
	if err != nil {
		slog.Error("Error loading config", "error", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
