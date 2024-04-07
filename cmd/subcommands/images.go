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

type imagesSubcommandOptions struct {
	isQuiet bool
}

type imagesSubcommandHandler struct {
	options *imagesSubcommandOptions
	client  client.Client
}

func NewImagesSubcommand(ctx context.Context, clientFactory factory.ClientFactory) *cobra.Command {
	isc := &imagesSubcommandHandler{
		options: &imagesSubcommandOptions{},
	}

	cmd := &cobra.Command{
		Use:   "images",
		Short: shortImagesDesc,
		Long:  longImagesDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Flags().Parse(args); err != nil {
				return err
			}

			// call completeHandler not to create client as many as the number of subcommands
			if err := isc.completeHandler(clientFactory); err != nil {
				return err
			}

			if err := isc.run(ctx); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVarP(&isc.options.isQuiet, "quiet", "q", false, "Only show image IDs")

	return cmd
}

func (ish *imagesSubcommandHandler) completeHandler(clientFactory factory.ClientFactory) error {
	var err error
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

	if !ish.options.isQuiet {
		fmt.Fprintln(writer, imageColumns)
	}

	for _, image := range imageList {
		/*
			hardcoding. not sure if a case of image.RepoTags size is 0.
			for now, just assume that case does not exists.
		*/
		imageID := strings.Split(image.ID, ":")[1][0:imageIDLength]
		repository := strings.Split(image.RepoTags[0], ":")[0]
		tag := strings.Split(image.RepoTags[0], ":")[1]
		size := humanize.Bytes(uint64(image.Size))
		created := time.Unix(image.Created, 0)

		if ish.options.isQuiet {
			fmt.Fprintf(writer, imageOutputQuietFormat, imageID)
		} else {
			fmt.Fprintf(writer, imageOutputFormat, repository, tag, imageID, humanize.Time(created), size)
		}
	}

	return nil
}
