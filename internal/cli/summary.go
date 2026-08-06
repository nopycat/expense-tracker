package cli

import (
	"expensetracker/internal/storage"
	"fmt"

	"github.com/spf13/cobra"
)

var monthMap = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func NewSummaryCmd(store *storage.Storage) *cobra.Command {
	var month int

	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Show expenses summary",
		Long: `Display the total amount of expenses.

Without the --month flag, shows the total expenses.
With the --month flag, shows expenses for a specific month of the current year.

Examples:
  expense-tracker summary
  expense-tracker summary --month 8`,
		Run: func(cmd *cobra.Command, args []string) {
			if month == 0 {
				total := store.Summary()
				fmt.Printf("Total expenses: $%.2f\n", total)
			} else if 1 <= month && month <= 12 {
				total := store.SummaryByMonth(month)
				fmt.Printf("Total expenses for %s: $%.2f\n", monthMap[month], total)
			} else {
				fmt.Println("Month must be between 1 and 12.")
			}
		},
	}

	cmd.Flags().IntVarP(
		&month,
		"month",
		"m",
		0,
		"show expenses for a specific month",
	)

	return cmd
}
