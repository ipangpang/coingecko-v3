package nfts

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetNFTsListEndpoint           = "/nfts/list"
	GetNFTDataEndpoint            = "/nfts/{id}"
	GetNFTContractDataEndpoint    = "/nfts/{asset_platform_id}/contract/{contract_address}"
	GetNFTsMarketDataEndpoint     = "/nfts/list/market_data"
	GetNFTHistoryEndpoint         = "/nfts/{id}/history"
	GetNFTContractHistoryEndpoint = "/nfts/{asset_platform_id}/contract/{contract_address}/history"
	GetNFTTickersEndpoint         = "/nfts/{id}/tickers"
)

type Client interface {
	GetNFTsList() (*GetNFTsListResponse, error)
	GetNFTData(request *GetNFTDataRequest) (*GetNFTDataResponse, error)
	GetNFTContractData(request *GetNFTContractDataRequest) (*GetNFTContractDataResponse, error)
	GetNFTsMarketData(request *GetNFTsMarketDataRequest) (*GetNFTsMarketDataResponse, error)
	GetNFTHistory(request *GetNFTHistoryRequest) (*GetNFTHistoryResponse, error)
	GetNFTContractHistory(request *GetNFTContractHistoryRequest) (*GetNFTContractHistoryResponse, error)
	GetNFTTickers(request *GetNFTTickersRequest) (*GetNFTTickersResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetNFTsList() (*GetNFTsListResponse, error) {
	var response GetNFTsListResponse

	if err := c.baseClient.Get(GetNFTsListEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetNFTData(request *GetNFTDataRequest) (*GetNFTDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetNFTDataResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
	}

	if err := c.baseClient.Get(GetNFTDataEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetNFTContractData(request *GetNFTContractDataRequest) (*GetNFTContractDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetNFTContractDataResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"asset_platform_id": request.AssetPlatformID,
			"contract_address":  request.ContractAddress,
		},
	}

	if err := c.baseClient.Get(GetNFTContractDataEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetNFTsMarketData(request *GetNFTsMarketDataRequest) (*GetNFTsMarketDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetNFTsMarketDataResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"order":     request.Order,
			"per_page":  fmt.Sprintf("%d", request.PerPage),
			"page":      fmt.Sprintf("%d", request.Page),
			"sparkline": fmt.Sprintf("%v", request.Sparkline),
		},
	}

	if err := c.baseClient.Get(GetNFTsMarketDataEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetNFTHistory(request *GetNFTHistoryRequest) (*GetNFTHistoryResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetNFTHistoryResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"days":        request.Days,
		},
	}

	if err := c.baseClient.Get(GetNFTHistoryEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetNFTContractHistory(request *GetNFTContractHistoryRequest) (*GetNFTContractHistoryResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetNFTContractHistoryResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"asset_platform_id": request.AssetPlatformID,
			"contract_address":  request.ContractAddress,
		},
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"days":        request.Days,
		},
	}

	if err := c.baseClient.Get(GetNFTContractHistoryEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetNFTTickers(request *GetNFTTickersRequest) (*GetNFTTickersResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetNFTTickersResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
		QueryParams: map[string]string{
			"exchange_ids":          request.ExchangeIDs,
			"include_exchange_logo": fmt.Sprintf("%v", request.IncludeExchangeLogo),
			"order":                 request.Order,
			"depth":                 fmt.Sprintf("%v", request.Depth),
		},
	}

	if err := c.baseClient.Get(GetNFTTickersEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
