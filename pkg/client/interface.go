package client

type Client interface {
	ListContainerImages() error
}
