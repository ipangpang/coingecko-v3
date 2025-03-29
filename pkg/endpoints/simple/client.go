package simple

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
	"strconv"
	"strings"
)

const (
	GetCoinPriceByIDsEndpoint          = "/simple/price"
	GetCoinPriceByTokenAddressEndpoint = "/simple/token_price/{id}"
	GetSupportedCurrenciesEndpoint     = "/simple/supported_vs_currencies"
)

type Client interface {
	GetCoinPriceByIDs(request *GetCoinPriceByIDsRequest) (*GetCoinPriceByIDsResponse, error)
	GetCoinPriceByTokenAddress(request *GetCoinPriceByTokenAddressRequest) (*GetCoinPriceByTokenAddressResponse, error)
	GetSupportedCurrencies() (*GetSupportedCurrenciesResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetCoinPriceByIDs(request *GetCoinPriceByIDsRequest) (*GetCoinPriceByIDsResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinPriceByIDsResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"ids":                     strings.Join(request.CoinIDs, ","),
			"vs_currencies":           strings.Join(request.VsCurrencies, ","),
			"include_market_cap":      strconv.FormatBool(request.IncludeMarketCap),
			"include_24hr_vol":        strconv.FormatBool(request.Include24HrVol),
			"include_24hr_change":     strconv.FormatBool(request.Include24HrChange),
			"include_last_updated_at": strconv.FormatBool(request.IncludeLastUpdatedAt),
			"precision":               request.Precision,
		},
	}

	if err := c.baseClient.Get(GetCoinPriceByIDsEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCoinPriceByTokenAddress(request *GetCoinPriceByTokenAddressRequest) (*GetCoinPriceByTokenAddressResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCoinPriceByTokenAddressResponse

	opts := &base.RequestOptions{
		PathParams: map[string]string{
			"id": request.ID,
		},
		QueryParams: map[string]string{
			"contract_addresses":      strings.Join(request.ContractAddresses, ","),
			"vs_currencies":           strings.Join(request.VsCurrencies, ","),
			"include_market_cap":      strconv.FormatBool(request.IncludeMarketCap),
			"include_24hr_vol":        strconv.FormatBool(request.Include24HrVol),
			"include_24hr_change":     strconv.FormatBool(request.Include24HrChange),
			"include_last_updated_at": strconv.FormatBool(request.IncludeLastUpdatedAt),
			"precision":               request.Precision,
		},
	}

	if err := c.baseClient.Get(GetCoinPriceByTokenAddressEndpoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetSupportedCurrencies() (*GetSupportedCurrenciesResponse, error) {
	var response GetSupportedCurrenciesResponse

	if err := c.baseClient.Get(GetSupportedCurrenciesEndpoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
