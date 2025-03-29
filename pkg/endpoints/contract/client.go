package contract

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetContractDataEndpoint             = "/coins/{asset_platform_id}/contract/{contract_address}"
	GetContractMarketChartEndpoint      = "/coins/{asset_platform_id}/contract/{contract_address}/market_chart"
	GetContractMarketChartRangeEndpoint = "/coins/{asset_platform_id}/contract/{contract_address}/market_chart/range"
)

type Client interface {
	GetContractData(request *GetContractDataRequest) (*GetContractDataResponse, error)
	GetContractMarketChart(request *GetContractMarketChartRequest) (*GetContractMarketChartResponse, error)
	GetContractMarketChartRange(request *GetContractMarketChartRangeRequest) (*GetContractMarketChartRangeResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetContractData(request *GetContractDataRequest) (*GetContractDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetContractDataResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"asset_platform_id": request.AssetPlatformID,
			"contract_address":  request.ContractAddress,
		},
	}

	if err := c.baseClient.Get(GetContractDataEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetContractMarketChart(request *GetContractMarketChartRequest) (*GetContractMarketChartResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetContractMarketChartResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"asset_platform_id": request.AssetPlatformID,
			"contract_address":  request.ContractAddress,
		},
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"days":        fmt.Sprintf("%d", request.Days),
		},
	}

	if err := c.baseClient.Get(GetContractMarketChartEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetContractMarketChartRange(request *GetContractMarketChartRangeRequest) (*GetContractMarketChartRangeResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetContractMarketChartRangeResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"asset_platform_id": request.AssetPlatformID,
			"contract_address":  request.ContractAddress,
		},
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"from":        fmt.Sprintf("%d", request.From),
			"to":          fmt.Sprintf("%d", request.To),
		},
	}

	if err := c.baseClient.Get(GetContractMarketChartRangeEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
