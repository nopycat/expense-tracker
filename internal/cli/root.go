package cli

import (
	"github.com/spf13/cobra"

	"expensetracker/internal/storage"
)

func NewRootCmd(storage *storage.Storage) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "Track and manage your expenses from the command line",
		Long: `ExpenseTracker is a command-line application for tracking personal expenses.

You can add, update, delete, list, and summarize expenses directly from your terminal.`,
	}

	rootCmd.AddCommand(
		NewAddCmd(storage),
		NewDeleteCmd(storage),
		NewListCmd(storage),
		NewSummaryCmd(storage),
		NewUpdateCmd(storage),
	)

	return rootCmd
}
