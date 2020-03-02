package main

import (
	"fmt"
	"os"

	"github.com/tsjohns9/gophercises/task-manager/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
