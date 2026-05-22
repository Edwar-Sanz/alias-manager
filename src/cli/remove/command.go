package remove

import (
	"am/src/functionality/amfile"
	"am/src/functionality/styles"
	"fmt"

	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "r [alias]",
		Aliases: []string{"remove", "rm", "delete", "del"},
		Short:   "Remove an alias",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dir, err := amfile.ConfigDir()
			if err != nil {
				fmt.Println(styles.ErrorStyle.Render("Error: " + err.Error()))
				return
			}
			filePath, err := amfile.CreateIfNotExist(dir)
			if err != nil {
				fmt.Println(styles.ErrorStyle.Render("Error: " + err.Error()))
				return
			}

			if err := amfile.DeleteAlias(filePath, args[0]); err != nil {
				fmt.Println(styles.ErrorStyle.Render("Error: " + err.Error()))
				return
			}

			fmt.Printf("alias %q removed\n", args[0])
		},
	}
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println(buildDescription())
	})
	return cmd
}
