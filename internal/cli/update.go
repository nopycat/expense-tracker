package cli

import (
	"expensetracker/internal/expense"
	"expensetracker/internal/storage"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

func NewUpdateCmd(store *storage.Storage) *cobra.Command {
	var description string
	var amount float64
	var date string
	var id int

	cmd := &cobra.Command{
		Use: "update",

		Short: "Update an existing expense",

		Long: `Update an existing expense in the tracker.

You can update the description, amount, date, or any combination of them.

Examples:
  expense-tracker update --id 1 --amount 25
  expense-tracker update --id 1 --description "Lunch"
  expense-tracker update --id 1 --date 2026-08-06`,
		Run: func(cmd *cobra.Command, args []string) {
			data := expense.UpdateData{}
			if cmd.Flags().Changed("description") {
				if strings.TrimSpace(description) == "" {
					fmt.Println("description cannot be empty")
					return
				}
				description = strings.TrimSpace(description)
				data.Description = &description

			}

			if cmd.Flags().Changed("amount") {
				if amount <= 0 {
					fmt.Println("amount must be positive")
					return
				}
				data.Amount = &amount
			}

			if cmd.Flags().Changed("date") {
				parsed, err := time.Parse("2006-01-02", date)
				if err != nil {
					fmt.Println("invalid date format, expected YYYY-MM-DD")
					return
				}
				data.Date = &parsed
			}

			if !cmd.Flags().Changed("amount") && !cmd.Flags().Changed("description") && !cmd.Flags().Changed("date") {
				fmt.Println("nothing to update")
				return
			}

			existing, err := store.Update(id, data)
			if err != nil {
				fmt.Println("id not found for update")
				return
			}

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			fmt.Fprintln(w, "ID\tDate\tDescription\tAmount")
			fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\n", existing.ID, existing.Date.Format("2006-01-02"), existing.Description, existing.Amount)

			w.Flush()

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

	cmd.Flags().StringVar(
		&date,
		"date",
		"",
		"expense date in YYYY-MM-DD format",
	)

	cmd.Flags().IntVar(
		&id,
		"id",
		0,
		"ID of the expense to update",
	)
	return cmd
}
