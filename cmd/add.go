package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"fmt"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
