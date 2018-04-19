package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// task := strings.Join(args, " ")
		fmt.Printf("I would be normally listing your tasks \n")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

