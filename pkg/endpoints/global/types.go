package global

// GetGlobalResponse represents the response from the Global API
type GetGlobalResponse struct {
	// Data contains the global data
	Data *GlobalData `json:"data"`
}

// GlobalData represents the global market data
type GlobalData struct {
	// ActiveCryptocurrencies is the number of active cryptocurrencies
	ActiveCryptocurrencies int `json:"active_cryptocurrencies"`
	// TotalMarketCap is the total market cap in different currencies
	TotalMarketCap map[string]float64 `json:"total_market_cap"`
	// TotalVolume is the total volume in different currencies
	TotalVolume map[string]float64 `json:"total_volume"`
	// MarketCapPercentage is the market cap percentage by coin
	MarketCapPercentage map[string]float64 `json:"market_cap_percentage"`
	// MarketCapChangePercentage24HUsd is the 24-hour market cap change percentage in USD
	MarketCapChangePercentage24HUsd float64 `json:"market_cap_change_percentage_24h_usd"`
	// UpdatedAt is the last update timestamp
	UpdatedAt int64 `json:"updated_at"`
}

// GetGlobalDefiResponse represents the response from the Global DeFi API
type GetGlobalDefiResponse struct {
	// Data contains the global DeFi data
	Data *GlobalDefiData `json:"data"`
}

// GlobalDefiData represents the global DeFi market data
type GlobalDefiData struct {
	// DefiMarketCap is the DeFi market cap in USD
	DefiMarketCap float64 `json:"defi_market_cap"`
	// EthMarketCap is the ETH market cap in USD
	EthMarketCap float64 `json:"eth_market_cap"`
	// DefiToEthRatio is the ratio of DeFi market cap to ETH market cap
	DefiToEthRatio float64 `json:"defi_to_eth_ratio"`
	// TradingVolume24H is the 24-hour trading volume in USD
	TradingVolume24H float64 `json:"trading_volume_24h"`
	// DefiDominance is the DeFi dominance percentage
	DefiDominance float64 `json:"defi_dominance"`
	// TopCoinName is the name of the top DeFi coin
	TopCoinName string `json:"top_coin_name"`
	// TopCoinDefiDominance is the dominance percentage of the top DeFi coin
	TopCoinDefiDominance float64 `json:"top_coin_defi_dominance"`
}

// GetGlobalMarketCapChartResponse represents the response from the Global Market Cap Chart API
type GetGlobalMarketCapChartResponse struct {
	// MarketCaps is a list of [timestamp, market_cap] pairs
	MarketCaps [][]float64 `json:"market_caps"`
	// TotalCaps is a list of [timestamp, total_cap] pairs
	TotalCaps [][]float64 `json:"total_caps"`
}
