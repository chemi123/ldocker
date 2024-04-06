package subcommands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewImagesSubCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "images",
		Short: shortImagesDesc,
		Long:  longImagesDesc,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("images called")
		},
	}
}
