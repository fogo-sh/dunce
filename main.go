package main

import (
	"fmt"
	"os"

	"github.com/fogo-sh/dunce/bot"
)

func main() {
	err := bot.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running Dunce: %s\n", err)
		os.Exit(1)
	}
}
