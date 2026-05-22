package list

import (
	"am/src/functionality/aliases"
	"am/src/functionality/amfile"
	"am/src/functionality/styles"
	"fmt"

	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	var category string
	var listCategories bool

	cmd := &cobra.Command{
		Use:     "l",
		Aliases: []string{"list", "ls"},
		Short:   "List all aliases",
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
			content, err := amfile.GetFileContent(filePath)
			if err != nil {
				fmt.Println(styles.ErrorStyle.Render("Error: " + err.Error()))
				return
			}

			parsed := aliases.ParseAliases(content)

			if listCategories {
				fmt.Println(aliases.FormatCategories(parsed.Categories))
				return
			}

			if category != "" {
				entries, ok := parsed.ByCategory[category]
				if !ok {
					fmt.Println(styles.ErrorStyle.Render("Error: category " + category + " not found"))
					return
				}
				fmt.Println(aliases.FormatCategory(category, entries))
				return
			}

			fmt.Println(aliases.FormatAllAliases(parsed))
		},
	}

	cmd.Flags().StringVarP(&category, "category", "c", "", "Filter by category")
	cmd.Flags().BoolVarP(&listCategories, "categories", "C", false, "List category names only")

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println(buildDescription())
	})
	return cmd
}
