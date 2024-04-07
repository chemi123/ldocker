package subcommands

import (
	"context"
	"fmt"

	"github.com/chemi123/ldocker/pkg/client"
	"github.com/chemi123/ldocker/pkg/factory"
	"github.com/spf13/cobra"
)

type ImagesSubcommandOptions struct{}

type imagesSubcommandHandler struct {
	options *ImagesSubcommandOptions
	client  client.Client
}

func NewImagesSubcommand(ctx context.Context, clientFactory factory.ClientFactory) *cobra.Command {
	isc := &imagesSubcommandHandler{}

	return &cobra.Command{
		Use:   "images",
		Short: shortImagesDesc,
		Long:  longImagesDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := isc.completeHandler(ctx, clientFactory); err != nil {
				return err
			}

			if err := isc.run(ctx); err != nil {
				return err
			}

			return nil
		},
	}
}

func (ish *imagesSubcommandHandler) completeOptions(ctx context.Context) error {
	return nil
}

func (ish *imagesSubcommandHandler) completeHandler(ctx context.Context, clientFactory factory.ClientFactory) error {
	err := ish.completeOptions(ctx)
	if err != nil {
		return err
	}

	ish.client, err = clientFactory.NewClient()
	if err != nil {
		return err
	}

	return nil
}

func (ish *imagesSubcommandHandler) run(ctx context.Context) error {
	imageList, err := ish.client.GetImageList(ctx)
	if err != nil {
		return err
	}

	for _, image := range imageList {
		fmt.Println(image)
	}

	return nil
}
