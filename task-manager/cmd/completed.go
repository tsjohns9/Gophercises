package cmd

import (
	"fmt"

	"../db"
	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Retrieve all completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.List()
		if err != nil {
			panic(err)
		}
		fmt.Println("All completed tasks:")
		for _, task := range tasks {
			if task.Completed == true {
				fmt.Printf("%d. %s\n", task.Key, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
