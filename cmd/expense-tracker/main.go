package main

import "expensetracker/internal/storage"

func main() {
	storage := storage.New()
	_ = storage
}
