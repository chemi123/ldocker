package subcommands

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/chemi123/ldocker/pkg/client"
	"github.com/chemi123/ldocker/pkg/factory"
	"github.com/dustin/go-humanize"
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

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
	defer writer.Flush()

	fmt.Fprintln(writer, imageColumns)

	for _, image := range imageList {
		repository := strings.Split(image.RepoTags[0], ":")[0]
		tag := strings.Split(image.RepoTags[0], ":")[1]
		imageID := strings.Split(image.ID, ":")[1][0:imageIDLength]
		size := humanize.Bytes(uint64(image.Size))
		created := time.Unix(image.Created, 0)
		fmt.Fprintf(writer, imageOutputFormat, repository, tag, imageID, humanize.Time(created), size)
	}

	return nil
}
