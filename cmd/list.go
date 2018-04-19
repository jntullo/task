package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		listTasks()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

