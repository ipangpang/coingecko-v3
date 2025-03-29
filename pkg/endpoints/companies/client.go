package companies

import (
	"fmt"

	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetPublicTreasuryEndpoint = "/companies/public_treasury/bitcoin"
)

type Client struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return Client{
		baseClient: baseClient,
	}
}

func (c *Client) GetPublicTreasury() (*GetPublicTreasuryResponse, error) {
	var response GetPublicTreasuryResponse

	err := c.baseClient.Get(GetPublicTreasuryEndpoint, &base.RequestOptions{
		QueryParams: map[string]string{
			"timeout": "10s",
		},
	}, &response)

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return &response, nil
}
