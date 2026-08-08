package storage

import (
	"errors"
	"expensetracker/internal/expense"
	"sort"
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

func NewFromData(data expense.FileData) *Storage {
	exs := make(map[int]expense.Expense, len(data.Expenses))
	for i := 0; i < len(data.Expenses); i++ {
		exs[data.Expenses[i].ID] = data.Expenses[i]
	}
	return &Storage{
		expenses: exs,
		nextID:   data.NextID,
	}
}

func (s *Storage) Export() expense.FileData {
	return expense.FileData{
		NextID:   s.nextID,
		Expenses: s.List(),
	}
}

func (s *Storage) Add(description string, amount float64, time time.Time) int {
	s.nextID++
	id := s.nextID
	s.expenses[id] = expense.Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Date:        time,
	}
	return id
}

func (s *Storage) Update(id int, data expense.UpdateData) (expense.Expense, error) {
	existing, ok := s.expenses[id]
	if !ok {
		return expense.Expense{}, errors.New("UpdateExpenseNotFound")
	}

	if data.Description != nil {
		existing.Description = *data.Description
	}
	if data.Amount != nil {
		existing.Amount = *data.Amount
	}
	if data.Date != nil {
		existing.Date = *data.Date
	}
	s.expenses[id] = existing
	return existing, nil
}

func (s *Storage) Delete(id int) error {
	if _, ok := s.expenses[id]; ok {
		delete(s.expenses, id)
		return nil
	}
	return errors.New("DeleteExpenseNotFound")
}

func (s *Storage) List() []expense.Expense {
	expenseList := make([]expense.Expense, 0, len(s.expenses))
	for _, v := range s.expenses {
		expenseList = append(expenseList, v)
	}
	sort.Slice(expenseList, func(i, j int) bool {
		return expenseList[i].ID < expenseList[j].ID
	})
	return expenseList
}

func (s *Storage) Summary() float64 {
	sum := 0.0
	for _, v := range s.expenses {
		sum += v.Amount
	}
	return sum
}

func (s *Storage) SummaryByMonth(month int) float64 {
	sum := 0.0
	currYear := time.Now().Year()
	for _, v := range s.expenses {
		if int(v.Date.Month()) == month && v.Date.Year() == currYear {
			sum += v.Amount
		}
	}
	return sum
}
