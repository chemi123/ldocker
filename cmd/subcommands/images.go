package subcommands

import (
	"github.com/chemi123/ldocker/pkg/client"
	"github.com/chemi123/ldocker/pkg/factory"
	"github.com/spf13/cobra"
)

type ImagesSubcommandOptions struct{}

type imagesSubcommandHandler struct {
	options *ImagesSubcommandOptions
	client  client.Client
}

func (ish *imagesSubcommandHandler) completeOptions() error {
	return nil
}

func (ish *imagesSubcommandHandler) completeHandler(clientFactory factory.ClientFactory) error {
	err := ish.completeOptions()
	if err != nil {
		return err
	}

	ish.client, err = clientFactory.NewClient()
	if err != nil {
		return err
	}

	return nil
}

func (ish *imagesSubcommandHandler) run() error {
	return ish.client.ListContainerImages()
}

func NewImagesSubCommand(clientFactory factory.ClientFactory) *cobra.Command {
	isc := &imagesSubcommandHandler{}

	return &cobra.Command{
		Use:   "images",
		Short: shortImagesDesc,
		Long:  longImagesDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := isc.completeHandler(clientFactory); err != nil {
				return err
			}

			if err := isc.run(); err != nil {
				return err
			}

			return nil
		},
	}
}
