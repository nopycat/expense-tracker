package main

import (
	"expensetracker/internal/cli"
	"expensetracker/internal/storage"
)

func main() {
	store := storage.New()

	rootCmd := cli.NewRootCmd(store)

	if err := rootCmd.Execute(); err != nil {
		return
	}
}
