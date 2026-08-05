package storage

import (
	"errors"
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

/*TODO
func (s *Storage) Update(id int) error {
	if _, ok := s.expenses[id]; ok {
		delete(s.expenses, id)
		return nil
	}
	return errors.New("expense not found")
}
*/

func (s *Storage) Delete(id int) error {
	if _, ok := s.expenses[id]; ok {
		delete(s.expenses, id)
		return nil
	}
	return errors.New("expense not found")
}
