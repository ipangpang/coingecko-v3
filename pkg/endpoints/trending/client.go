package trending

import (
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetTrendingEndpoint = "/search/trending"
)

type Client interface {
	GetTrending() (*TrendingResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetTrending() (*TrendingResponse, error) {
	var response TrendingResponse

	if err := c.baseClient.Get(GetTrendingEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
