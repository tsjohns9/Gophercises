package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI task manager with Cobra",
}
