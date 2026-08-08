package json

import (
	"encoding/json"
	"expensetracker/internal/expense"
	"expensetracker/internal/storage"
	"io"
	"os"
)

func Load(path string) (*storage.Storage, error) {
	file, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data expense.FileData
	err = decoder.Decode(&data)
	if err == io.EOF {
		return storage.New(), nil
	}

	if err != nil {
		return nil, err
	}

	return storage.NewFromData(data), nil
}

func Save(path string, store *storage.Storage) error {
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	data := store.Export()
	return encoder.Encode(data)
}
