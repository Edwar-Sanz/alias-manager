package add

import (
	"am/src/constants"
	"am/src/functionality/amfile"
	"am/src/functionality/sanitize"
	"am/src/functionality/styles"
	"am/src/types"
	"fmt"

	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "a [alias] [command] [category] [description]",
		Aliases: []string{"add"},
		Short:   "Add or update an alias",
		Args:    cobra.RangeArgs(2, 4),
		Run: func(cmd *cobra.Command, args []string) {
			category, desc := constants.UnCategorizedCategory, ""
			if len(args) >= 3 {
				category = args[2]
			}
			if len(args) == 4 {
				desc = args[3]
			}

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

			entry := sanitize.Entry(types.AliasEntry{Name: args[0], Command: args[1], Category: category, Desc: desc})
			if err := amfile.WriteAlias(filePath, entry); err != nil {
				fmt.Println(styles.ErrorStyle.Render("Error: " + err.Error()))
				return
			}

			fmt.Printf("alias %q saved\n", args[0])
		},
	}
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println(buildDescription())
	})
	return cmd
}
