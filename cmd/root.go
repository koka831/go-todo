package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"

	"github.com/koka831/go-todo/model"
)

var filepath string := "~/todo-go.json"

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo cli tool",
	Long:  "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		// show the list of TODO
		show()
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().StringVarP(&filepath, "config", "c", "", "config file")
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("todofile not found. initializing...")

	}
}

func show() {
	fmt.Println("show")
}
