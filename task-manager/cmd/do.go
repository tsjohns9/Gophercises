package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"../db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Execute a task",
	Run: func(cmd *cobra.Command, args []string) {
		taskNumString := strings.Join(args, " ")
		taskNum, _ := strconv.Atoi(taskNumString)
		if err := db.Delete(taskNum); err != nil {
			panic(err)
		}
		fmt.Println("Task completed")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
