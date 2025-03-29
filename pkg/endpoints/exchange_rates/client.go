package exchange_rates

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetExchangeRatesRequestPoint = "/exchange_rates"
	GetExchangeRateEndpoint      = "/exchange_rates/{id}"
)

type Client interface {
	GetExchangeRates() (*GetExchangeRatesResponse, error)
	GetExchangeRate(request *GetExchangeRateRequest) (*GetExchangeRateResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetExchangeRates() (*GetExchangeRatesResponse, error) {
	var response GetExchangeRatesResponse

	if err := c.baseClient.Get(GetExchangeRatesRequestPoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetExchangeRate(request *GetExchangeRateRequest) (*GetExchangeRateResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetExchangeRateResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
	}

	if err := c.baseClient.Get(GetExchangeRateEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
