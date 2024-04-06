package factory

import "github.com/chemi123/ldocker/pkg/client"

type ClientFactory interface {
	NewClient() (client.Client, error)
}
