package expense

import (
	"fmt"
	"strings"
	"time"
)

type Expense struct {
	ID          uint32    `json:"id"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Date        time.Time `json:"time"`
}

const filePath = "expenses.json"

func AddExpense(amount float64, category, description string) error {
	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	expense := Expense{
		ID:          uint32(len(expenses) + 1),
		Amount:      amount,
		Category:    category,
		Description: description,
		Date:        time.Now(),
	}

	expenses = append(expenses, expense)
	return saveExpenses(expenses)
}

func ViewExpenses(category string) error {
	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	formatString := "%-10v%-22v%-16v%-20v%-40v\n"
	fmt.Printf(formatString, "ID", "Date", "Amount", "Category", "Description")
	for _, expense := range expenses {
		fmt.Printf("%-10v%-22v$%-15.2f%-20v%-40v\n", expense.ID, expense.Date.Format("01/02/2006 3:04 PM"), expense.Amount,
			expense.Category, expense.Description)
	}

	return nil
}

func DeleteExpense(id uint32) error {
	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	deleted := false
	for i := range expenses {
		if expenses[i].ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			deleted = true
			break
		}
	}

	if !deleted {
		return fmt.Errorf("expense ID %v not found", id)
	}

	for i := range expenses {
		expenses[i].ID = uint32(i + 1)
	}

	saveExpenses(expenses)

	return nil
}

func GenerateReport(start, end, category string) error {
	var startDate, endDate time.Time
	var err error
	if start != "" {
		startDate, err = ParseDate(start)
		if err != nil {
			return err
		}
	}
	if end != "" {
		endDate, err = ParseDate(end)
		if err != nil {
			return err
		}
		endDate = endDate.Add(24 * time.Hour)
	}

	fmt.Println("Expense Report")
	fmt.Println("--------------")

	expenses, err := loadExpenses()
	if err != nil {
		return err
	}

	totalAmount := 0.0
	count := 0
	formatString := "%-10v%-22v%-20v%-20v%-20v\n"
	fmt.Printf(formatString, "ID", "Date", "Amount", "Category", "Description")
	for _, expense := range expenses {
		if !startDate.IsZero() && expense.Date.Before(startDate) {
			continue
		}
		if !endDate.IsZero() && expense.Date.After(endDate) {
			continue
		}
		if category != "" && !strings.EqualFold(expense.Category, category) {
			continue
		}
		fmt.Printf("%-10v%-22v$%-19.2f%-20v%-20v\n", expense.ID, expense.Date.Format("01/02/2006 3:04 PM"), expense.Amount,
			expense.Category, expense.Description)
		totalAmount += expense.Amount
		count++
	}

	fmt.Println("--------------")
	fmt.Printf("Total Expenses: %d\n", count)
	fmt.Printf("Total Amount Spent: $%.2f\n", totalAmount)

	return nil
}

func ParseDate(date string) (time.Time, error) {
	return time.ParseInLocation("01/02/2006", date, time.Local)
}
