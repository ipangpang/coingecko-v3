package search

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	SearchEndpoint = "/search"
)

type Client interface {
	Search(request *SearchRequest) (*SearchResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) Search(request *SearchRequest) (*SearchResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response SearchResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"query": request.Query,
		},
	}

	if err := c.baseClient.Get(SearchEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
