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
		t := Task{ Name: task, Complete: false }
		err := t.add()
		if err != nil {
			fmt.Printf("ERROR! %s", err)
		} else {
			fmt.Printf("Added \"%s\" to your task list.\n", task)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
