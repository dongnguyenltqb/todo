package cmd

import (
	"todo/pkg/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "this command list all task in 3 list : todo,doing and done ",
	Run: func(cmd *cobra.Command, args []string) {
		todo.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
