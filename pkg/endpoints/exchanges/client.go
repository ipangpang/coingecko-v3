package exchanges

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetExchangesListRequestPoint       = "/exchanges"
	GetExchangesListIDRequestPoint     = "/exchanges/list"
	GetExchangeDataRequestPoint        = "/exchanges/{id}"
	GetExchangeTickersRequestPoint     = "/exchanges/{id}/tickers"
	GetExchangeVolumeChartRequestPoint = "/exchanges/{id}/volume_chart"
)

type Client interface {
	GetExchangesList(request *GetExchangesListRequest) (*GetExchangesListResponse, error)
	GetExchangesListID() (*GetExchangesListIDResponse, error)
	GetExchangeData(request *GetExchangeDataRequest) (*GetExchangeDataResponse, error)
	GetExchangeTickers(request *GetExchangeTickersRequest) (*GetExchangeTickersResponse, error)
	GetExchangeVolumeChart(request *GetExchangeVolumeChartRequest) (*GetExchangeVolumeChartResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetExchangesList(request *GetExchangesListRequest) (*GetExchangesListResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetExchangesListResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"per_page": fmt.Sprintf("%d", request.PerPage),
			"page":     fmt.Sprintf("%d", request.Page),
		},
	}

	if err := c.baseClient.Get(GetExchangesListRequestPoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetExchangesListID() (*GetExchangesListIDResponse, error) {
	var response GetExchangesListIDResponse

	if err := c.baseClient.Get(GetExchangesListIDRequestPoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetExchangeData(request *GetExchangeDataRequest) (*GetExchangeDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetExchangeDataResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
	}

	if err := c.baseClient.Get(GetExchangeDataRequestPoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetExchangeTickers(request *GetExchangeTickersRequest) (*GetExchangeTickersResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetExchangeTickersResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
		QueryParams: map[string]string{
			"coin_ids":              fmt.Sprintf("%v", request.CoinIDs),
			"include_exchange_logo": fmt.Sprintf("%v", request.IncludeExchangeLogo),
			"page":                  fmt.Sprintf("%d", request.Page),
			"depth":                 fmt.Sprintf("%v", request.Depth),
			"order":                 request.Order,
		},
	}

	if err := c.baseClient.Get(GetExchangeTickersRequestPoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetExchangeVolumeChart(request *GetExchangeVolumeChartRequest) (*GetExchangeVolumeChartResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetExchangeVolumeChartResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
		QueryParams: map[string]string{
			"days": fmt.Sprintf("%d", request.Days),
		},
	}

	if err := c.baseClient.Get(GetExchangeVolumeChartRequestPoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
