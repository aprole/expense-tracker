/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/aprole/expense-tracker/pkg/expense"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Long:  `Add a new expense to the tracker`,
	Run: func(cmd *cobra.Command, args []string) {
		amount, _ := cmd.Flags().GetFloat64("amount")
		category, _ := cmd.Flags().GetString("category")
		description, _ := cmd.Flags().GetString("description")

		err := expense.AddExpense(amount, category, description)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding expense:%v\n", err)
			return
		}

		fmt.Println("Expense added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().Float64P("amount", "a", 0.0, "Amount of the expense")
	addCmd.Flags().StringP("category", "c", "", "Category of the expense")
	addCmd.Flags().StringP("description", "d", "", "Description of the expense")

	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagRequired("category")
}
