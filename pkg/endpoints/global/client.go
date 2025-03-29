package global

import (
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetGlobalRequestPoint               = "/global"
	GetGlobalDefiRequestPoint           = "/global/decentralized_finance_defi"
	GetGlobalMarketCapChartRequestPoint = "/global/market_cap_chart"
)

type Client interface {
	GetGlobal() (*GetGlobalResponse, error)
	GetGlobalDefi() (*GetGlobalDefiResponse, error)
	GetGlobalMarketCapChart(vsCurrency string, days string) (*GetGlobalMarketCapChartResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetGlobal() (*GetGlobalResponse, error) {
	var response GetGlobalResponse

	if err := c.baseClient.Get(GetGlobalRequestPoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetGlobalDefi() (*GetGlobalDefiResponse, error) {
	var response GetGlobalDefiResponse

	if err := c.baseClient.Get(GetGlobalDefiRequestPoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetGlobalMarketCapChart(vsCurrency string, days string) (*GetGlobalMarketCapChartResponse, error) {
	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"vs_currency": vsCurrency,
			"days":        days,
		},
	}

	var response GetGlobalMarketCapChartResponse
	if err := c.baseClient.Get(GetGlobalMarketCapChartRequestPoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
