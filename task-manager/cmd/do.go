package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tsjohns9/gophercises/task-manager/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		taskNumString := strings.Join(args, " ")
		taskNum, err := strconv.Atoi(taskNumString)
		if err != nil {
			fmt.Printf("A number is required to complete a task. Received %s\n", taskNumString)
			return
		}
		if err = db.Complete(taskNum); err != nil {
			panic(err)
		}
		fmt.Println("Task marked as completed")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
