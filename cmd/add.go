package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/koka831/go-todo/model"
)

init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd &cobra.Command {
	Use:	"add",
	Short:	"add new todo",
	Long:	"add new todo",
	Run:	func(cmd *cobra.Command, args []string) {
	
	}
}
