package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"strconv"
	"fmt"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := strconv.Atoi(task)
		fmt.Printf("The id - %d \n", id)
		if err != nil {
			fmt.Println("Error!")
		}
		doTask(id)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

