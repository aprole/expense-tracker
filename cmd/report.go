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

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a summary report of expenses",
	Long: `Generate a detailed summary report of your expenses.
You can filter the report by date range or category.
For example:

expense-tracker report --start 05/01/2024 --end 05/31/2024 --category Groceries`,
	Run: func(cmd *cobra.Command, args []string) {
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		category, _ := cmd.Flags().GetString("category")

		if err := expense.GenerateReport(start, end, category); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating report:%v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)

	reportCmd.Flags().StringP("start", "s", "", "Start date of the report (DD/MM/YYYY)")
	reportCmd.Flags().StringP("end", "e", "", "End date of the report (DD/MM/YYYY)")
	reportCmd.Flags().StringP("category", "c", "", "Category of the report")
}
