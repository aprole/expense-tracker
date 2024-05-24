/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aprole/expense-tracker/pkg/expense"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Long: `Delete an expense from the tracker with a given ID. 
For example:

expense-tracker delete 5`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID:%v\n", err)
			return
		}

		err = expense.DeleteExpense(uint32(id))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting expense: %v\n", err)
			return
		}

		fmt.Println("Expense deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// addCmd.Flags().StringP("start")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
