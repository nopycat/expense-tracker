package expense

import "time"

type Expense struct {
	ID          int
	Date        time.Duration
	Description string
	Amount      int
}
