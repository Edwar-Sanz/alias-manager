package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "am",
		Short:        "Alias Manager",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(BuildLongDescription())
		},
	}

	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println(BuildLongDescription())
	})
	return cmd
}
