package commands

import (
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-todo",
	Short: "go-todo is a simple cli that implements todo list",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("there was an error trying to start cli")
		os.Exit(1)
	}
}
