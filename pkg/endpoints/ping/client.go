package ping

import (
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	PingEndpoint = "/ping"
)

type Client interface {
	Ping() (*PingResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) Ping() (*PingResponse, error) {
	var response PingResponse

	if err := c.baseClient.Get(PingEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
