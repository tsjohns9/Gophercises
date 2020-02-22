package cmd

import (
	"fmt"
	"strings"

	"../db"
	"github.com/spf13/cobra"
)

// Task is a task
type Task struct {
	Key   int
	Value string
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if err := db.Create(task); err != nil {
			panic(err)
		}
		fmt.Println("Task added")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
