package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "Task",
	Short: "A simple task manager cli tool",
}
