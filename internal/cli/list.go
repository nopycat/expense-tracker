package cli

import (
	"expensetracker/internal/storage"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func NewListCmd(store *storage.Storage) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		Long: `Display all expenses stored in the tracker.

Shows ID, date, description and amount for each expense.`,
		Run: func(cmd *cobra.Command, args []string) {
			expenseList := store.List()

			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

			fmt.Fprintln(w, "ID\tDate\tDescription\tAmount")
			for _, item := range expenseList {
				fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\n", item.ID, item.Date.Format("2006-01-02"), item.Description, item.Amount)
			}

			w.Flush()

		},
	}

	return cmd
}
