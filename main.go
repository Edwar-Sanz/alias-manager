package main

import (
	"am/src/cli/add"
	list "am/src/cli/list"
	"am/src/cli/remove"
	root "am/src/cli/root"
	"am/src/functionality/styles"
	"fmt"
)

func main() {
	rootCmd := root.CreateCommand()
	rootCmd.AddCommand(add.CreateCommand())
	rootCmd.AddCommand(list.CreateCommand())
	rootCmd.AddCommand(remove.CreateCommand())

	cmd, err := rootCmd.ExecuteC()
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error: " + err.Error()))
		fmt.Println()
		cmd.Help()
	}
}
