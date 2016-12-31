package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command {
	Use:	"version",
	Short:	"print the version of todo-go",
	Long:	"this is todo-go's",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("todo-go ver 0 --dev")
	},
}
