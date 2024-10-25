package main

import (
	"log/slog"
	"os"

	"github.com/fogo-sh/dunce/cmd"
)

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	slog.SetDefault(slog.New(handler))

	cmd.Execute()
}
