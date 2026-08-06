package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"expensetracker/internal/storage"
)

func NewAddCmd(store *storage.Storage) *cobra.Command {
	var description string
	var amount float64

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new expense",
		Run: func(cmd *cobra.Command, args []string) {

			if strings.TrimSpace(description) == "" {
				fmt.Println("description cannot be empty")
				return
			}

			if amount <= 0 {
				fmt.Println("amount must be positive")
				return
			}

			id := store.Add(description, amount, time.Now())

			fmt.Printf("Expense added successfully (ID: %d)\n", id)
		},
	}

	cmd.Flags().StringVarP(
		&description,
		"description",
		"d",
		"",
		"expense description",
	)

	cmd.Flags().Float64VarP(
		&amount,
		"amount",
		"a",
		0,
		"expense amount",
	)

	return cmd
}
