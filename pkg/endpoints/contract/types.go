package contract

import "github.com/go-playground/validator/v10"

// GetContractDataRequest represents the request parameters for getting contract data
type GetContractDataRequest struct {
	// AssetPlatformID is the ID of the asset platform
	AssetPlatformID string `json:"asset_platform_id" validate:"required"`
	// ContractAddress is the contract address
	ContractAddress string `json:"contract_address" validate:"required"`
}

// GetContractDataResponse represents the response from the Contract Data API
type GetContractDataResponse struct {
	// ID is the unique identifier of the coin
	ID string `json:"id"`
	// Symbol is the symbol of the coin
	Symbol string `json:"symbol"`
	// Name is the name of the coin
	Name string `json:"name"`
	// AssetPlatformID is the ID of the asset platform
	AssetPlatformID string `json:"asset_platform_id"`
	// ContractAddress is the contract address
	ContractAddress string `json:"contract_address"`
	// MarketData is the market data of the coin
	MarketData MarketData `json:"market_data"`
}

// MarketData represents market data for a coin
type MarketData struct {
	// CurrentPrice is the current price of the coin
	CurrentPrice map[string]float64 `json:"current_price"`
	// TotalValueLocked is the total value locked
	TotalValueLocked map[string]float64 `json:"total_value_locked,omitempty"`
	// MarketCap is the market capitalization
	MarketCap map[string]float64 `json:"market_cap"`
	// MarketCapRank is the rank of the coin by market cap
	MarketCapRank int `json:"market_cap_rank"`
	// FullyDilutedValuation is the fully diluted valuation
	FullyDilutedValuation map[string]float64 `json:"fully_diluted_valuation,omitempty"`
	// TotalVolume is the total volume
	TotalVolume map[string]float64 `json:"total_volume"`
	// High24H is the 24-hour high price
	High24H map[string]float64 `json:"high_24h"`
	// Low24H is the 24-hour low price
	Low24H map[string]float64 `json:"low_24h"`
	// PriceChange24H is the 24-hour price change
	PriceChange24H float64 `json:"price_change_24h"`
	// PriceChangePercentage24H is the 24-hour price change percentage
	PriceChangePercentage24H float64 `json:"price_change_percentage_24h"`
	// MarketCapChange24H is the 24-hour market cap change
	MarketCapChange24H float64 `json:"market_cap_change_24h"`
	// MarketCapChangePercentage24H is the 24-hour market cap change percentage
	MarketCapChangePercentage24H float64 `json:"market_cap_change_percentage_24h"`
	// CirculatingSupply is the circulating supply
	CirculatingSupply float64 `json:"circulating_supply"`
	// TotalSupply is the total supply
	TotalSupply float64 `json:"total_supply"`
	// MaxSupply is the maximum supply
	MaxSupply float64 `json:"max_supply,omitempty"`
	// ATH is the all-time high price
	ATH map[string]float64 `json:"ath"`
	// ATHChangePercentage is the all-time high price change percentage
	ATHChangePercentage map[string]float64 `json:"ath_change_percentage"`
	// ATHDate is the all-time high date
	ATHDate map[string]string `json:"ath_date"`
	// ATL is the all-time low price
	ATL map[string]float64 `json:"atl"`
	// ATLChangePercentage is the all-time low price change percentage
	ATLChangePercentage map[string]float64 `json:"atl_change_percentage"`
	// ATLDate is the all-time low date
	ATLDate map[string]string `json:"atl_date"`
}

// GetContractMarketChartRequest represents the request parameters for getting contract market chart data
type GetContractMarketChartRequest struct {
	// AssetPlatformID is the ID of the asset platform
	AssetPlatformID string `json:"asset_platform_id" validate:"required"`
	// ContractAddress is the contract address
	ContractAddress string `json:"contract_address" validate:"required"`
	// VsCurrency is the target currency
	VsCurrency string `json:"vs_currency" validate:"required"`
	// Days is the number of days of data to return
	Days int `json:"days" validate:"required,min=1,max=365"`
}

// GetContractMarketChartResponse represents the response from the Contract Market Chart API
type GetContractMarketChartResponse [][]interface{}

// GetContractMarketChartRangeRequest represents the request parameters for getting contract market chart data with time range
type GetContractMarketChartRangeRequest struct {
	// AssetPlatformID is the ID of the asset platform
	AssetPlatformID string `json:"asset_platform_id" validate:"required"`
	// ContractAddress is the contract address
	ContractAddress string `json:"contract_address" validate:"required"`
	// VsCurrency is the target currency
	VsCurrency string `json:"vs_currency" validate:"required"`
	// From is the start date in Unix timestamp
	From int64 `json:"from" validate:"required"`
	// To is the end date in Unix timestamp
	To int64 `json:"to" validate:"required"`
}

// GetContractMarketChartRangeResponse represents the response from the Contract Market Chart Range API
type GetContractMarketChartRangeResponse [][]interface{}

// Validate validates the request parameters
func (r *GetContractDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetContractMarketChartRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetContractMarketChartRangeRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
