package cmd

import "log/slog"

func checkError(err error, message string) {
	if err != nil {
		slog.Error(message)
		panic(err)
	}
}
