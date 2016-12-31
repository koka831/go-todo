package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"

	"../model"
)

var filepath string

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
