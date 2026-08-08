package main

import (
	"expensetracker/internal/cli"
	js "expensetracker/internal/repository/json"
)

func main() {
	//store := storage.New()
	store, err := js.Load("data.json")
	if err != nil {
		return
	}
	rootCmd := cli.NewRootCmd(store)

	if err := rootCmd.Execute(); err != nil {
		return
	}
	if err := js.Save("data.json", store); err != nil {
		return
	}
}
