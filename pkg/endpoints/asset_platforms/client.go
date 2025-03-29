package asset_platforms

import (
	"fmt"

	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetAssetPlatformsEndpoint              = "/asset_platforms"
	GetTokenListsByAssetPlatformIDEndpoint = "/asset_platforms/{asset_platform_id}/contract"
)

type Client struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return Client{
		baseClient: baseClient,
	}
}

func (c *Client) GetAssetPlatforms() (*GetAssetPlatformsResponse, error) {
	var response GetAssetPlatformsResponse

	err := c.baseClient.Get(GetAssetPlatformsEndpoint, &base.RequestOptions{
		QueryParams: map[string]string{
			"timeout": "10s",
		},
	}, &response)

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return &response, nil
}

func (c *Client) GetTokenListsByAssetPlatformID(request *GetTokenListsByAssetPlatformIDRequest) (*GetTokenListsByAssetPlatformIDResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetTokenListsByAssetPlatformIDResponse

	err := c.baseClient.Get(GetTokenListsByAssetPlatformIDEndpoint, &base.RequestOptions{
		PathParams: map[string]string{
			"asset_platform_id": request.AssetPlatformID,
		},
		QueryParams: map[string]string{
			"timeout": "10s",
		},
	}, &response)

	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	return &response, nil
}
