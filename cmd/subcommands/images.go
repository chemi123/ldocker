package subcommands

import (
	"fmt"
	"os"

	"github.com/chemi123/ldocker/pkg/client"
	"github.com/chemi123/ldocker/pkg/factory"
	"github.com/spf13/cobra"
)

type ImagesSubcommandOptions struct{}

type imagesSubcommandHandler struct {
	options *ImagesSubcommandOptions
	client  client.Client
}

func (isc *imagesSubcommandHandler) completeOptions() error {
	return nil
}

func (isc *imagesSubcommandHandler) completeHandler(clientFactory factory.ClientFactory) error {
	err := isc.completeOptions()
	if err != nil {
		return err
	}

	isc.client, err = clientFactory.NewClient()
	if err != nil {
		return err
	}

	return nil
}

func (isc *imagesSubcommandHandler) run() error {
	return isc.client.ListContainerImages()
}

func NewImagesSubCommand(clientFactory factory.ClientFactory) *cobra.Command {
	isc := &imagesSubcommandHandler{}

	return &cobra.Command{
		Use:   "images",
		Short: shortImagesDesc,
		Long:  longImagesDesc,
		Run: func(cmd *cobra.Command, args []string) {
			if err := isc.completeHandler(clientFactory); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}

			if err := isc.run(); err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		},
	}
}
