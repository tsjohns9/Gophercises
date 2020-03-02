package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tsjohns9/gophercises/task-manager/db"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task from the list",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		task := strings.Join(args, " ")
		taskNum, err := strconv.Atoi(task)
		if err != nil {
			fmt.Printf("A number is required to delete a task. Received %s\n", task)
			return
		}
		if err = db.Delete(taskNum); err != nil {
			panic(err)
		}
		fmt.Println("Task deleted")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
