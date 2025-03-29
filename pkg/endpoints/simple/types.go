package simple

import "github.com/go-playground/validator/v10"

// GetCoinPriceByIDsRequest represents the request parameters for getting coin prices by IDs
type GetCoinPriceByIDsRequest struct {
	// CoinIDs is a list of coin IDs to get prices for
	CoinIDs []string `json:"ids" validate:"required,min=1"`
	// VsCurrencies is a list of target currencies to get prices in (e.g., ["usd", "eur"])
	VsCurrencies []string `json:"vs_currencies" validate:"required,min=1"`
	// IncludeMarketCap indicates whether to include market cap data
	IncludeMarketCap bool `json:"include_market_cap,omitempty"`
	// Include24HrVol indicates whether to include 24hr volume data
	Include24HrVol bool `json:"include_24hr_vol,omitempty"`
	// Include24HrChange indicates whether to include 24hr price change data
	Include24HrChange bool `json:"include_24hr_change,omitempty"`
	// IncludeLastUpdatedAt indicates whether to include last updated timestamp
	IncludeLastUpdatedAt bool `json:"include_last_updated_at,omitempty"`
	// Precision indicates the number of decimal places to use for price data
	Precision string `json:"precision,omitempty"`
}

// CoinPrice represents the price data for a single coin
type CoinPrice struct {
	// USD price
	USD float64 `json:"usd,omitempty"`
	// USD market cap
	USDMarketCap float64 `json:"usd_market_cap,omitempty"`
	// USD 24h volume
	USD24HrVol float64 `json:"usd_24h_vol,omitempty"`
	// USD 24h price change percentage
	USD24HrChange float64 `json:"usd_24h_change,omitempty"`
	// Last updated timestamp
	LastUpdatedAt int64 `json:"last_updated_at,omitempty"`
}

// GetCoinPriceByIDsResponse represents the response from the Simple Price API
type GetCoinPriceByIDsResponse map[string]CoinPrice

// GetCoinPriceByTokenAddressRequest represents the request parameters for getting coin price by token address
type GetCoinPriceByTokenAddressRequest struct {
	// ID is the token contract address
	ID string `json:"id" validate:"required"`
	// ContractAddresses is a list of token contract addresses
	ContractAddresses []string `json:"contract_addresses" validate:"required,min=1"`
	// VsCurrencies is a list of target currencies to get prices in (e.g., ["usd", "eur"])
	VsCurrencies []string `json:"vs_currencies" validate:"required,min=1"`
	// IncludeMarketCap indicates whether to include market cap data
	IncludeMarketCap bool `json:"include_market_cap,omitempty"`
	// Include24HrVol indicates whether to include 24hr volume data
	Include24HrVol bool `json:"include_24hr_vol,omitempty"`
	// Include24HrChange indicates whether to include 24hr price change data
	Include24HrChange bool `json:"include_24hr_change,omitempty"`
	// IncludeLastUpdatedAt indicates whether to include last updated timestamp
	IncludeLastUpdatedAt bool `json:"include_last_updated_at,omitempty"`
	// Precision indicates the number of decimal places to use for price data
	Precision string `json:"precision,omitempty"`
}

// CoinPriceByTokenAddress represents the price data for a single token
type CoinPriceByTokenAddress struct {
	// USD price
	USD float64 `json:"usd,omitempty"`
	// USD market cap
	USDMarketCap float64 `json:"usd_market_cap,omitempty"`
	// USD 24h volume
	USD24HrVol float64 `json:"usd_24h_vol,omitempty"`
	// USD 24h price change percentage
	USD24HrChange float64 `json:"usd_24h_change,omitempty"`
	// Last updated timestamp
	LastUpdatedAt int64 `json:"last_updated_at,omitempty"`
}

// GetCoinPriceByTokenAddressResponse represents the response from the Simple Token Price API
type GetCoinPriceByTokenAddressResponse map[string]CoinPriceByTokenAddress

// GetSupportedCurrenciesResponse represents the response from the Supported Currencies API
type GetSupportedCurrenciesResponse []string

// Validate validates the request structs
func (r *GetCoinPriceByIDsRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinPriceByTokenAddressRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
