package client

import (
	"context"

	"github.com/docker/docker/api/types/image"
)

type Client interface {
	GetImageList(context.Context) ([]image.Summary, error)
}
