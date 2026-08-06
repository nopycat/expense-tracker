package cli

import (
	"expensetracker/internal/storage"
	"fmt"

	"github.com/spf13/cobra"
)

func NewDeleteCmd(storage *storage.Storage) *cobra.Command {
	var id int

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an expense by ID",
		Long: `Delete an existing expense from the tracker.

You need to provide the ID of the expense you want to delete.
Example:
  expense-tracker delete --id 2`,
		Run: func(cmd *cobra.Command, args []string) {
			err := storage.Delete(id)
			if err != nil {
				fmt.Println("id not found")
				return
			}
			fmt.Println("Expense deleted successfully")
		},
	}
	cmd.Flags().IntVar(
		&id,
		"id",
		0,
		"expense id",
	)
	return cmd
}
