package key

import (
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	KeyEndpoint = "/key"
)

type Client interface {
	Key() (*KeyResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) Key() (*KeyResponse, error) {
	var response KeyResponse

	if err := c.baseClient.Get(KeyEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
