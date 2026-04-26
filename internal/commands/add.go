package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a todo to a list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
