package coins

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
	"strings"
)

const (
	GetCoinsListEndpoint                       = "/coins/list"
	GetTopGainersAndLosersEndpoint             = "/coins/top_gainers_losers"
	GetRecentlyAddedCoinsEndpoint              = "/coins/list/new"
	GetCoinsListWithMarketDataEndpoint         = "/coins/markets"
	GetCoinDataByIDEndpoint                    = "/coins/{id}"
	GetCoinTickersByIDEndpoint                 = "/coins/{id}/tickers"
	GetCoinHistoryByIDEndpoint                 = "/coins/{id}/history"
	GetCoinMarketChartByIDEndpoint             = "/coins/{id}/market_chart"
	GetCoinMarketChartRangeEndpoint            = "/coins/{id}/market_chart/range"
	GetCoinOHLCByIDEndpoint                    = "/coins/{id}/ohlc"
	GetCoinOHLCRangeEndpoint                   = "/coins/{id}/ohlc/range"
	GetCoinCirculatingSupplyChartEndpoint      = "/coins/{id}/circulating_supply_chart"
	GetCoinCirculatingSupplyChartRangeEndpoint = "/coins/{id}/circulating_supply_chart/range"
	GetCoinTotalSupplyChartEndpoint            = "/coins/{id}/total_supply_chart"
	GetCoinTotalSupplyChartRangeEndpoint       = "/coins/{id}/total_supply_chart/range"
)

type Client interface {
	GetCoinsList(request *GetCoinsListRequest) (*GetCoinsListResponse, error)
	GetTopGainersAndLosers(request *GetTopGainersAndLosersRequest) (*GetTopGainersAndLosersResponse, error)
	GetRecentlyAddedCoins() (*GetRecentlyAddedCoinsResponse, error)
	GetCoinsListWithMarketData(request *GetCoinsListWithMarketDataRequest) (*GetCoinsListWithMarketDataResponse, error)
	GetCoinDataByID(request *GetCoinDataByIDRequest) (*GetCoinDataByIDResponse, error)
	GetCoinTickersByID(request *GetCoinTickersByIDRequest) (*GetCoinTickersByIDResponse, error)
	GetCoinHistoryByID(request *GetCoinHistoryByIDRequest) (*GetCoinHistoryByIDResponse, error)
	GetCoinMarketChartByID(request *GetCoinMarketChartByIDRequest) (*GetCoinMarketChartByIDResponse, error)
	GetCoinMarketChartRange(request *GetCoinMarketChartRangeRequest) (*GetCoinMarketChartRangeResponse, error)
	GetCoinOHLCByID(request *GetCoinOHLCByIDRequest) (*GetCoinOHLCByIDResponse, error)
	GetCoinOHLCRange(request *GetCoinOHLCRangeRequest) (*GetCoinOHLCRangeResponse, error)
	GetCoinCirculatingSupplyChart(request *GetCoinCirculatingSupplyChartRequest) (*GetCoinCirculatingSupplyChartResponse, error)
	GetCoinCirculatingSupplyChartRange(request *GetCoinCirculatingSupplyChartRangeRequest) (*GetCoinCirculatingSupplyChartRangeResponse, error)
	GetCoinTotalSupplyChart(request *GetCoinTotalSupplyChartRequest) (*GetCoinTotalSupplyChartResponse, error)
	GetCoinTotalSupplyChartRange(request *GetCoinTotalSupplyChartRangeRequest) (*GetCoinTotalSupplyChartRangeResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetCoinsList(request *GetCoinsListRequest) (*GetCoinsListResponse, error) {
	var response GetCoinsListResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"include_platform": fmt.Sprintf("%v", request.IncludePlatform),
			"status":           request.Status,
		},
	}

	if err := c.baseClient.Get(GetCoinsListEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetTopGainersAndLosers(request *GetTopGainersAndLosersRequest) (*GetTopGainersAndLosersResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetTopGainersAndLosersResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"duration":    request.Duration,
			"top_coins":   fmt.Sprintf("%d", request.TopCoins),
		},
	}

	if err := c.baseClient.Get(GetTopGainersAndLosersEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetRecentlyAddedCoins() (*GetRecentlyAddedCoinsResponse, error) {
	var response GetRecentlyAddedCoinsResponse

	if err := c.baseClient.Get(GetRecentlyAddedCoinsEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinsListWithMarketData(request *GetCoinsListWithMarketDataRequest) (*GetCoinsListWithMarketDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinsListWithMarketDataResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"vs_currency":             request.VsCurrency,
			"ids":                     strings.Join(request.IDs, ","),
			"category":                request.Category,
			"order":                   request.Order,
			"per_page":                fmt.Sprintf("%d", request.PerPage),
			"page":                    fmt.Sprintf("%d", request.Page),
			"sparkline":               fmt.Sprintf("%v", request.Sparkline),
			"price_change_percentage": strings.Join(request.PriceChangePercentage, ","),
			"locale":                  request.Locale,
		},
	}

	if err := c.baseClient.Get(GetCoinsListWithMarketDataEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinDataByID(request *GetCoinDataByIDRequest) (*GetCoinDataByIDResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinDataByIDResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
		QueryParams: map[string]string{
			"localization":   fmt.Sprintf("%v", request.Localization),
			"tickers":        fmt.Sprintf("%v", request.Tickers),
			"market_data":    fmt.Sprintf("%v", request.MarketData),
			"community_data": fmt.Sprintf("%v", request.CommunityData),
			"developer_data": fmt.Sprintf("%v", request.DeveloperData),
			"sparkline":      fmt.Sprintf("%v", request.Sparkline),
			"locale":         request.Locale,
		},
	}

	if err := c.baseClient.Get(GetCoinDataByIDEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinTickersByID(request *GetCoinTickersByIDRequest) (*GetCoinTickersByIDResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinTickersByIDResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"exchange_ids": strings.Join(request.ExchangeIDs, ","),
			"include":      strings.Join(request.Include, ","),
			"order":        request.Order,
			"depth":        fmt.Sprintf("%v", request.Depth),
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinTickersByIDEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinHistoryByID(request *GetCoinHistoryByIDRequest) (*GetCoinHistoryByIDResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinHistoryByIDResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"date":         request.Date,
			"localization": fmt.Sprintf("%v", request.Localization),
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinHistoryByIDEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinMarketChartByID(request *GetCoinMarketChartByIDRequest) (*GetCoinMarketChartByIDResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinMarketChartByIDResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"days":        request.Days,
			"interval":    request.Interval,
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinMarketChartByIDEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinMarketChartRange(request *GetCoinMarketChartRangeRequest) (*GetCoinMarketChartRangeResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinMarketChartRangeResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"from":        fmt.Sprintf("%d", request.From),
			"to":          fmt.Sprintf("%d", request.To),
			"interval":    request.Interval,
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinMarketChartRangeEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinOHLCByID(request *GetCoinOHLCByIDRequest) (*GetCoinOHLCByIDResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinOHLCByIDResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"days":        request.Days,
			"interval":    request.Interval,
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinOHLCByIDEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinOHLCRange(request *GetCoinOHLCRangeRequest) (*GetCoinOHLCRangeResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinOHLCRangeResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"vs_currency": request.VsCurrency,
			"from":        fmt.Sprintf("%d", request.From),
			"to":          fmt.Sprintf("%d", request.To),
			"interval":    request.Interval,
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinOHLCRangeEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinCirculatingSupplyChart(request *GetCoinCirculatingSupplyChartRequest) (*GetCoinCirculatingSupplyChartResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinCirculatingSupplyChartResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"days": request.Days,
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinCirculatingSupplyChartEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinCirculatingSupplyChartRange(request *GetCoinCirculatingSupplyChartRangeRequest) (*GetCoinCirculatingSupplyChartRangeResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinCirculatingSupplyChartRangeResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"from": fmt.Sprintf("%d", request.From),
			"to":   fmt.Sprintf("%d", request.To),
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinCirculatingSupplyChartRangeEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinTotalSupplyChart(request *GetCoinTotalSupplyChartRequest) (*GetCoinTotalSupplyChartResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinTotalSupplyChartResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"days": request.Days,
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinTotalSupplyChartEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinTotalSupplyChartRange(request *GetCoinTotalSupplyChartRangeRequest) (*GetCoinTotalSupplyChartRangeResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinTotalSupplyChartRangeResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"from": fmt.Sprintf("%d", request.From),
			"to":   fmt.Sprintf("%d", request.To),
		},
	}

	if err := c.baseClient.Get(strings.Replace(GetCoinTotalSupplyChartRangeEndpoint, "{id}", request.ID, -1), opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
