package storage

import (
	"expensetracker/internal/expense"
	"time"
)

type Storage struct {
	expenses map[int]expense.Expense
	nextID   int
}

func New() *Storage {
	return &Storage{
		expenses: make(map[int]expense.Expense),
		nextID:   0,
	}
}

func (s *Storage) Add(description string, amount float64, time time.Time) int {
	s.nextID++
	id := s.nextID
	s.expenses[id] = expense.Expense{
		Description: description,
		Amount:      amount,
		Date:        time,
	}
	return id
}
