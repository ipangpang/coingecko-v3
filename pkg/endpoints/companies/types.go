package companies

// GetPublicTreasuryResponse represents the response from the Public Treasury API
type GetPublicTreasuryResponse struct {
	// TotalHoldings is the total holdings in different currencies
	TotalHoldings map[string]float64 `json:"total_holdings"`
	// TotalValueUsd is the total value in USD
	TotalValueUsd float64 `json:"total_value_usd"`
	// MarketCapDominance is the market cap dominance percentage
	MarketCapDominance float64 `json:"market_cap_dominance"`
	// Companies contains the list of companies
	Companies []Company `json:"companies"`
}

// Company represents a single company
type Company struct {
	// Name is the name of the company
	Name string `json:"name"`
	// Symbol is the stock symbol of the company
	Symbol string `json:"symbol"`
	// Country is the country where the company is located
	Country string `json:"country"`
	// TotalHoldings is the total holdings in different currencies
	TotalHoldings map[string]float64 `json:"total_holdings"`
	// TotalEntryValueUsd is the total entry value in USD
	TotalEntryValueUsd float64 `json:"total_entry_value_usd"`
	// TotalCurrentValueUsd is the current total value in USD
	TotalCurrentValueUsd float64 `json:"total_current_value_usd"`
	// PercentageOfTotalSupply is the percentage of total supply held
	PercentageOfTotalSupply float64 `json:"percentage_of_total_supply"`
}
