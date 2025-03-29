package derivatives

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetDerivativesListEndpoint            = "/derivatives"
	GetDerivativesExchangesListEndpoint   = "/derivatives/exchanges"
	GetDerivativeExchangeDataEndpoint     = "/derivatives/exchanges/{id}"
	GetDerivativesExchangesListIDEndpoint = "/derivatives/exchanges/list"
)

type Client interface {
	GetDerivativesList() (*GetDerivativesListResponse, error)
	GetDerivativesExchangesList() (*GetDerivativesExchangesListResponse, error)
	GetDerivativeExchangeData(request *GetDerivativeExchangeDataRequest) (*GetDerivativeExchangeDataResponse, error)
	GetDerivativesExchangesListIDMap() (*GetDerivativesExchangesListIDMapResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetDerivativesList() (*GetDerivativesListResponse, error) {
	var response GetDerivativesListResponse

	if err := c.baseClient.Get(GetDerivativesListEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetDerivativesExchangesList() (*GetDerivativesExchangesListResponse, error) {
	var response GetDerivativesExchangesListResponse

	if err := c.baseClient.Get(GetDerivativesExchangesListEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetDerivativeExchangeData(request *GetDerivativeExchangeDataRequest) (*GetDerivativeExchangeDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetDerivativeExchangeDataResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
		QueryParams: map[string]string{
			"include_tickers": fmt.Sprintf("%v", request.IncludeTickers),
		},
	}

	if err := c.baseClient.Get(GetDerivativeExchangeDataEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetDerivativesExchangesListIDMap() (*GetDerivativesExchangesListIDMapResponse, error) {
	var response GetDerivativesExchangesListIDMapResponse

	if err := c.baseClient.Get(GetDerivativesExchangesListIDEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
