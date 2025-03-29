package derivatives

import "github.com/go-playground/validator/v10"

// Derivative represents a single derivative
type Derivative struct {
	// Market is the market name
	Market string `json:"market"`
	// Symbol is the symbol of the derivative
	Symbol string `json:"symbol"`
	// IndexID is the index ID
	IndexID string `json:"index_id"`
	// Price is the price
	Price float64 `json:"price"`
	// PriceChangePercentage24H is the 24-hour price change percentage
	PriceChangePercentage24H float64 `json:"price_change_percentage_24h"`
	// ContractType is the type of contract
	ContractType string `json:"contract_type"`
	// MarkPrice is the mark price
	MarkPrice float64 `json:"mark_price,omitempty"`
	// IndexPrice is the index price
	IndexPrice float64 `json:"index_price,omitempty"`
	// FundingRate is the funding rate
	FundingRate float64 `json:"funding_rate,omitempty"`
	// OpenInterest is the open interest
	OpenInterest float64 `json:"open_interest,omitempty"`
	// Volume24H is the 24-hour volume
	Volume24H float64 `json:"volume_24h,omitempty"`
	// LastUpdated is the last update timestamp
	LastUpdated string `json:"last_updated"`
}

// GetDerivativesListResponse represents the response from the Derivatives List API
type GetDerivativesListResponse []Derivative

// GetDerivativesExchangesListResponse represents the response from the Derivatives Exchanges List API
type GetDerivativesExchangesListResponse []DerivativeExchange

// DerivativeExchange represents a single derivative exchange
type DerivativeExchange struct {
	// Name is the name of the exchange
	Name string `json:"name"`
	// ID is the unique identifier of the exchange
	ID string `json:"id"`
	// OpenInterestBTC is the open interest in BTC
	OpenInterestBTC float64 `json:"open_interest_btc"`
	// TradeVolume24HBTC is the 24-hour trading volume in BTC
	TradeVolume24HBTC float64 `json:"trade_volume_24h_btc"`
	// NumberOfPerpetualPairs is the number of perpetual pairs
	NumberOfPerpetualPairs int `json:"number_of_perpetual_pairs"`
	// NumberOfFuturesPairs is the number of futures pairs
	NumberOfFuturesPairs int `json:"number_of_futures_pairs"`
	// Image is the URL of the exchange's logo
	Image string `json:"image"`
	// YearEstablished is the year the exchange was established
	YearEstablished int `json:"year_established,omitempty"`
	// Country is the country where the exchange is located
	Country string `json:"country,omitempty"`
	// Description is the description of the exchange
	Description string `json:"description,omitempty"`
	// URL is the website URL of the exchange
	URL string `json:"url,omitempty"`
}

// GetDerivativeExchangeDataRequest represents the request parameters for getting derivative exchange data
type GetDerivativeExchangeDataRequest struct {
	// ID is the unique identifier of the exchange
	ID string `json:"id" validate:"required"`
	// IncludeTickers is whether to include tickers
	IncludeTickers string `json:"include_tickers,omitempty" validate:"omitempty,oneof=unexpired all"`
}

// GetDerivativeExchangeDataResponse represents the response from the Derivative Exchange Data API
type GetDerivativeExchangeDataResponse struct {
	// Name is the name of the exchange
	Name string `json:"name"`
	// ID is the unique identifier of the exchange
	ID string `json:"id"`
	// OpenInterestBTC is the open interest in BTC
	OpenInterestBTC float64 `json:"open_interest_btc"`
	// TradeVolume24HBTC is the 24-hour trading volume in BTC
	TradeVolume24HBTC float64 `json:"trade_volume_24h_btc"`
	// NumberOfPerpetualPairs is the number of perpetual pairs
	NumberOfPerpetualPairs int `json:"number_of_perpetual_pairs"`
	// NumberOfFuturesPairs is the number of futures pairs
	NumberOfFuturesPairs int `json:"number_of_futures_pairs"`
	// Image is the URL of the exchange's logo
	Image string `json:"image"`
	// YearEstablished is the year the exchange was established
	YearEstablished int `json:"year_established,omitempty"`
	// Country is the country where the exchange is located
	Country string `json:"country,omitempty"`
	// Description is the description of the exchange
	Description string `json:"description,omitempty"`
	// URL is the website URL of the exchange
	URL string `json:"url,omitempty"`
	// Tickers is the list of tickers
	Tickers []DerivativeTicker `json:"tickers,omitempty"`
}

// DerivativeTicker represents a single derivative ticker
type DerivativeTicker struct {
	// Symbol is the symbol of the derivative
	Symbol string `json:"symbol"`
	// Base is the base currency
	Base string `json:"base"`
	// Target is the target currency
	Target string `json:"target"`
	// Market is the market name
	Market string `json:"market"`
	// Last is the last price
	Last float64 `json:"last"`
	// Volume is the volume
	Volume float64 `json:"volume"`
	// ConvertedLast is the converted last price
	ConvertedLast map[string]float64 `json:"converted_last,omitempty"`
	// ConvertedVolume is the converted volume
	ConvertedVolume map[string]float64 `json:"converted_volume,omitempty"`
	// TrustScore is the trust score
	TrustScore string `json:"trust_score,omitempty"`
	// BidAskSpreadPercentage is the bid-ask spread percentage
	BidAskSpreadPercentage float64 `json:"bid_ask_spread_percentage,omitempty"`
	// Timestamp is the timestamp
	Timestamp string `json:"timestamp,omitempty"`
	// LastTradedAt is the last traded at timestamp
	LastTradedAt string `json:"last_traded_at,omitempty"`
	// LastFetchAt is the last fetch at timestamp
	LastFetchAt string `json:"last_fetch_at,omitempty"`
	// IsAnomaly indicates whether the ticker is anomalous
	IsAnomaly bool `json:"is_anomaly,omitempty"`
	// IsStale indicates whether the ticker is stale
	IsStale bool `json:"is_stale,omitempty"`
	// TradeURL is the trade URL
	TradeURL string `json:"trade_url,omitempty"`
	// TokenInfoURL is the token info URL
	TokenInfoURL string `json:"token_info_url,omitempty"`
	// CoinID is the coin ID
	CoinID string `json:"coin_id,omitempty"`
	// TargetCoinID is the target coin ID
	TargetCoinID string `json:"target_coin_id,omitempty"`
	// FundingRate is the funding rate
	FundingRate float64 `json:"funding_rate,omitempty"`
	// OpenInterest is the open interest
	OpenInterest float64 `json:"open_interest,omitempty"`
	// NextFundingRate is the next funding rate
	NextFundingRate float64 `json:"next_funding_rate,omitempty"`
	// ContractType is the type of contract
	ContractType string `json:"contract_type,omitempty"`
	// ExpirationTimestamp is the expiration timestamp
	ExpirationTimestamp string `json:"expiration_timestamp,omitempty"`
}

// GetDerivativesExchangesListIDMapResponse represents the response from the Derivatives Exchanges List (ID Map) API
type GetDerivativesExchangesListIDMapResponse map[string]DerivativeExchange

// Validate validates the request parameters
func (r *GetDerivativeExchangeDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
