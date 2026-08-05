package expense

import "time"

type Expense struct {
	ID          int
	Date        time.Time
	Description string
	Amount      float64
}

type UpdateData struct {
	Description *string
	Amount      *float64
	Date        *time.Time
}
