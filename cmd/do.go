package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"fmt"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Marked \"%s\" as completed.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

