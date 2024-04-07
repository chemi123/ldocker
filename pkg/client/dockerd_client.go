package client

import (
	"context"

	"github.com/docker/docker/api/types/image"
	dclient "github.com/docker/docker/client"
)

type DockerdClient struct {
	client *dclient.Client
}

func (dc *DockerdClient) GetImageList(ctx context.Context) ([]image.Summary, error) {
	imageList, err := dc.client.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return nil, err
	}
	return imageList, nil
}

func NewDockerdClient() (Client, error) {
	dc, err := dclient.NewClientWithOpts(dclient.FromEnv, dclient.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return &DockerdClient{client: dc}, nil
}
