package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tsjohns9/gophercises/task-manager/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.List()
		if err != nil {
			panic(err)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete!")
			return
		}
		fmt.Println("You have the following tasks:")
		for _, task := range tasks {
			fmt.Printf("%d. %s - Completed: %t \n", task.Key, task.Value, task.Completed)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
