package main

import (
	"expensetracker/internal/cli"
	"expensetracker/internal/storage"
)

func main() {
	storage := storage.New()

	rootCmd := cli.NewRootCmd(storage)

	if err := rootCmd.Execute(); err != nil {
		return
	}
}
