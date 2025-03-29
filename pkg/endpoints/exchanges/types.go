package exchanges

import "github.com/go-playground/validator/v10"

// GetExchangesListRequest represents the request parameters for getting exchanges list
type GetExchangesListRequest struct {
	// PerPage is the number of results per page
	PerPage int `json:"per_page,omitempty" validate:"omitempty,min=1,max=250"`
	// Page is the page number
	Page int `json:"page,omitempty" validate:"omitempty,min=1"`
}

// GetExchangesListResponse represents the response from the Exchanges List API
type GetExchangesListResponse []Exchange

// Exchange represents a single exchange
type Exchange struct {
	// ID is the unique identifier of the exchange
	ID string `json:"id"`
	// Name is the name of the exchange
	Name string `json:"name"`
	// YearEstablished is the year the exchange was established
	YearEstablished int `json:"year_established,omitempty"`
	// Country is the country where the exchange is located
	Country string `json:"country,omitempty"`
	// Description is the description of the exchange
	Description string `json:"description,omitempty"`
	// URL is the website URL of the exchange
	URL string `json:"url,omitempty"`
	// Image is the URL of the exchange's logo
	Image string `json:"image,omitempty"`
	// HasTradingIncentive indicates whether the exchange has trading incentives
	HasTradingIncentive bool `json:"has_trading_incentive,omitempty"`
	// TrustScore is the trust score of the exchange
	TrustScore int `json:"trust_score,omitempty"`
	// TrustScoreRank is the rank of the exchange by trust score
	TrustScoreRank int `json:"trust_score_rank,omitempty"`
	// TradeVolume24HBtc is the 24-hour trading volume in BTC
	TradeVolume24HBtc float64 `json:"trade_volume_24h_btc,omitempty"`
	// TradeVolume24HBtcNormalized is the normalized 24-hour trading volume in BTC
	TradeVolume24HBtcNormalized float64 `json:"trade_volume_24h_btc_normalized,omitempty"`
}

// GetExchangesListIDResponse represents the response from the Exchanges List (ID Map) API
type GetExchangesListIDResponse map[string]Exchange

// GetExchangeDataRequest represents the request parameters for getting exchange data
type GetExchangeDataRequest struct {
	// ID is the unique identifier of the exchange
	ID string `json:"id" validate:"required"`
}

// GetExchangeDataResponse represents the response from the Exchange Data API
type GetExchangeDataResponse struct {
	// Name is the name of the exchange
	Name string `json:"name"`
	// YearEstablished is the year the exchange was established
	YearEstablished int `json:"year_established,omitempty"`
	// Country is the country where the exchange is located
	Country string `json:"country,omitempty"`
	// Description is the description of the exchange
	Description string `json:"description,omitempty"`
	// URL is the website URL of the exchange
	URL string `json:"url,omitempty"`
	// Image is the URL of the exchange's logo
	Image string `json:"image,omitempty"`
	// FacebookURL is the Facebook URL of the exchange
	FacebookURL string `json:"facebook_url,omitempty"`
	// RedditURL is the Reddit URL of the exchange
	RedditURL string `json:"reddit_url,omitempty"`
	// TelegramURL is the Telegram URL of the exchange
	TelegramURL string `json:"telegram_url,omitempty"`
	// SlackURL is the Slack URL of the exchange
	SlackURL string `json:"slack_url,omitempty"`
	// OtherURL1 is the first other URL of the exchange
	OtherURL1 string `json:"other_url_1,omitempty"`
	// OtherURL2 is the second other URL of the exchange
	OtherURL2 string `json:"other_url_2,omitempty"`
	// TwitterHandle is the Twitter handle of the exchange
	TwitterHandle string `json:"twitter_handle,omitempty"`
	// HasTradingIncentive indicates whether the exchange has trading incentives
	HasTradingIncentive bool `json:"has_trading_incentive,omitempty"`
	// Centralized indicates whether the exchange is centralized
	Centralized bool `json:"centralized,omitempty"`
	// PublicNotice is the public notice of the exchange
	PublicNotice string `json:"public_notice,omitempty"`
	// AML indicates whether the exchange has AML
	AML bool `json:"aml,omitempty"`
	// KYC indicates whether the exchange has KYC
	KYC bool `json:"kyc,omitempty"`
	// Whitepaper is the URL of the exchange's whitepaper
	Whitepaper string `json:"whitepaper,omitempty"`
	// StatusUpdates is the list of status updates
	StatusUpdates []string `json:"status_updates,omitempty"`
	// Links is the list of links
	Links map[string]string `json:"links,omitempty"`
	// TrustScore is the trust score of the exchange
	TrustScore int `json:"trust_score,omitempty"`
	// TrustScoreRank is the rank of the exchange by trust score
	TrustScoreRank int `json:"trust_score_rank,omitempty"`
	// TradeVolume24HBtc is the 24-hour trading volume in BTC
	TradeVolume24HBtc float64 `json:"trade_volume_24h_btc,omitempty"`
	// TradeVolume24HBtcNormalized is the normalized 24-hour trading volume in BTC
	TradeVolume24HBtcNormalized float64 `json:"trade_volume_24h_btc_normalized,omitempty"`
	// Tickers is the list of tickers
	Tickers []Ticker `json:"tickers,omitempty"`
}

// Ticker represents a single ticker
type Ticker struct {
	// Base is the base currency
	Base string `json:"base"`
	// Target is the target currency
	Target string `json:"target"`
	// Market is the market information
	Market Market `json:"market"`
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
}

// Market represents market information
type Market struct {
	// Identifier is the market identifier
	Identifier string `json:"identifier"`
	// Name is the market name
	Name string `json:"name"`
	// HasTradingIncentive indicates whether the market has trading incentives
	HasTradingIncentive bool `json:"has_trading_incentive,omitempty"`
}

// GetExchangeTickersRequest represents the request parameters for getting exchange tickers
type GetExchangeTickersRequest struct {
	// ID is the unique identifier of the exchange
	ID string `json:"id" validate:"required"`
	// CoinIDs is the list of coin IDs to filter by
	CoinIDs []string `json:"coin_ids,omitempty"`
	// IncludeExchangeLogo is whether to include exchange logo
	IncludeExchangeLogo bool `json:"include_exchange_logo,omitempty"`
	// Page is the page number
	Page int `json:"page,omitempty" validate:"omitempty,min=1"`
	// Depth is the depth of the order book
	Depth bool `json:"depth,omitempty"`
	// Order is the order to sort by
	Order string `json:"order,omitempty" validate:"omitempty,oneof=trust_score_desc trust_score_asc volume_desc volume_asc id_desc id_asc"`
}

// GetExchangeTickersResponse represents the response from the Exchange Tickers API
type GetExchangeTickersResponse struct {
	// Name is the name of the exchange
	Name string `json:"name"`
	// Tickers is the list of tickers
	Tickers []Ticker `json:"tickers"`
}

// GetExchangeVolumeChartRequest represents the request parameters for getting exchange volume chart
type GetExchangeVolumeChartRequest struct {
	// ID is the unique identifier of the exchange
	ID string `json:"id" validate:"required"`
	// Days is the number of days of data to return
	Days string `json:"days" validate:"required"`
}

// GetExchangeVolumeChartResponse represents the response from the Exchange Volume Chart API
type GetExchangeVolumeChartResponse [][]interface{}

// Validate validates the request parameters
func (r *GetExchangesListRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetExchangeDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetExchangeTickersRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetExchangeVolumeChartRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
