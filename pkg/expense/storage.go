package expense

import (
	"encoding/json"
	"os"
)

func loadExpenses() ([]Expense, error) {
	var expenses []Expense

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return expenses, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func saveExpenses(expenses []Expense) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(expenses)
}
