package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"todo/pkg/todo"
	"todo/pkg/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command add new entry for list",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			utils.ClearScreen()
			todoText := utils.ApplyStyle("bold", "yellow", "todo")
			doingText := utils.ApplyStyle("bold", "red", "doing")
			doneText := utils.ApplyStyle("bold", "green", "done")
			fmt.Println(`Enter "exit" to exit app`)
			fmt.Println("========================")
			fmt.Println("Please enter List Name: ", todoText, " || ", doingText, " || ", doneText)
			list := bufio.NewScanner(os.Stdin)
			list.Scan()
			if list.Text() == "exit" {
				os.Exit(0)
			}
			if list.Text() == "" {
				continue
			}
			fmt.Println("Please enter value")
			value := bufio.NewScanner(os.Stdin)
			value.Scan()
			if value.Text() == "exit" {
				os.Exit(0)
			}
			if value.Text() == "" {
				continue
			}
			err := todo.Add(list.Text(), value.Text())
			if err != nil {
				panic(err)
			}
			fmt.Println("=> done")
			time.Sleep(3 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
