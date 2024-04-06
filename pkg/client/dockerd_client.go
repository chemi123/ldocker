package client

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
	dclient "github.com/docker/docker/client"
)

type DockerdClient struct {
	client *dclient.Client
}

func (dc *DockerdClient) ListContainerImages() error {
	ctx := context.Background()

	images, err := dc.client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return err
	}

	for _, image := range images {
		fmt.Println(image)
	}

	return nil
}

func NewDockerdClient() (Client, error) {
	dc, err := dclient.NewClientWithOpts(dclient.FromEnv, dclient.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &DockerdClient{client: dc}, nil
}
