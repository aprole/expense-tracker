/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/aprole/expense-tracker/pkg/expense"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View expenses",
	Long:  `View expenses from the tracker`,
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		expense.ViewExpenses(category)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	viewCmd.Flags().StringP("category", "c", "", "Category of the expense")
}
