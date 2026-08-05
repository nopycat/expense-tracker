package storage

import (
	"expensetracker/internal/expense"
	"sync"
)

type storage struct {
	expenses []expense.Expense
	mx       *sync.Mutex
}

func New() *storage {
	return &storage{
		mx: &sync.Mutex{},
	}
}
