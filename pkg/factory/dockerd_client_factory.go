package factory

import (
	"github.com/chemi123/ldocker/pkg/client"
)

type DockerdClientFactory struct{}

func (dcf *DockerdClientFactory) NewClient() (client.Client, error) {
	return client.NewDockerdClient()
}

func NewDockerdClientFactory() *DockerdClientFactory {
	return &DockerdClientFactory{}
}
