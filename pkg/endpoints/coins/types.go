package coins

import "github.com/go-playground/validator/v10"

// GetCoinsListRequest represents the request parameters for getting coins list
type GetCoinsListRequest struct {
	// IncludePlatform indicates whether to include platform contract addresses
	IncludePlatform bool `json:"include_platform,omitempty"`
	// Status indicates whether to include inactive coins (active or inactive)
	Status string `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}

// Coin represents a single coin in the list
type Coin struct {
	// ID is the unique identifier of the coin
	ID string `json:"id"`
	// Symbol is the symbol of the coin
	Symbol string `json:"symbol"`
	// Name is the name of the coin
	Name string `json:"name"`
	// Platforms contains contract addresses on different platforms
	Platforms map[string]string `json:"platforms,omitempty"`
}

// GetCoinsListResponse represents the response from the Coins List API
type GetCoinsListResponse []Coin

// GetTopGainersAndLosersRequest represents the request parameters for getting top gainers and losers
type GetTopGainersAndLosersRequest struct {
	// VsCurrency is the target currency for market data
	VsCurrency string `json:"vs_currency" validate:"required"`
	// Duration is the time duration for price change calculation (1h, 24h, 7d, 14d, 30d, 200d, 1y)
	Duration string `json:"duration" validate:"required,oneof=1h 24h 7d 14d 30d 200d 1y"`
	// TopCoins is the market cap ranking filter (300-1000)
	TopCoins int `json:"top_coins,omitempty" validate:"omitempty,min=300,max=1000"`
}

// GetTopGainersAndLosersResponse represents the response from the Top Gainers & Losers API
type GetTopGainersAndLosersResponse struct {
	// TopGainers contains the top gaining coins
	TopGainers []CoinMarketData `json:"top_gainers"`
	// TopLosers contains the top losing coins
	TopLosers []CoinMarketData `json:"top_losers"`
}

// CoinMarketData represents market data for a coin
type CoinMarketData struct {
	// ID is the unique identifier of the coin
	ID string `json:"id"`
	// Symbol is the symbol of the coin
	Symbol string `json:"symbol"`
	// Name is the name of the coin
	Name string `json:"name"`
	// Image is the URL of the coin's image
	Image string `json:"image"`
	// CurrentPrice is the current price of the coin
	CurrentPrice float64 `json:"current_price"`
	// MarketCap is the market capitalization of the coin
	MarketCap float64 `json:"market_cap"`
	// MarketCapRank is the rank of the coin by market cap
	MarketCapRank int `json:"market_cap_rank"`
	// PriceChangePercentage24H is the 24-hour price change percentage
	PriceChangePercentage24H float64 `json:"price_change_percentage_24h"`
}

// GetRecentlyAddedCoinsResponse represents the response from the Recently Added Coins API
type GetRecentlyAddedCoinsResponse []CoinMarketData

// GetCoinsListWithMarketDataRequest represents the request parameters for getting coins list with market data
type GetCoinsListWithMarketDataRequest struct {
	// VsCurrency is the target currency for market data
	VsCurrency string `json:"vs_currency" validate:"required"`
	// IDs is a list of coin IDs to get data for
	IDs []string `json:"ids,omitempty"`
	// Category is the category to filter by
	Category string `json:"category,omitempty"`
	// Order is the order to sort the results by
	Order string `json:"order,omitempty" validate:"omitempty,oneof=market_cap_desc market_cap_asc gecko_desc gecko_asc volume_desc volume_asc id_desc id_asc"`
	// PerPage is the number of results per page
	PerPage int `json:"per_page,omitempty" validate:"omitempty,min=1,max=250"`
	// Page is the page number
	Page int `json:"page,omitempty" validate:"omitempty,min=1"`
	// Sparkline indicates whether to include sparkline data
	Sparkline bool `json:"sparkline,omitempty"`
	// PriceChangePercentage is the price change percentage to include
	PriceChangePercentage []string `json:"price_change_percentage,omitempty" validate:"omitempty,dive,oneof=1h 24h 7d 14d 30d 200d 1y"`
	// Locale is the language to use for localization
	Locale string `json:"locale,omitempty" validate:"omitempty,oneof=en de es fr it pt ru ko ja zh"`
}

// GetCoinsListWithMarketDataResponse represents the response from the Coins List with Market Data API
type GetCoinsListWithMarketDataResponse []CoinMarketData

// GetCoinDataByIDRequest represents the request parameters for getting coin data by ID
type GetCoinDataByIDRequest struct {
	// ID is the unique identifier of the coin
	ID string `json:"id" validate:"required"`
	// Localization indicates whether to include localization data
	Localization bool `json:"localization,omitempty"`
	// Tickers indicates whether to include tickers data
	Tickers bool `json:"tickers,omitempty"`
	// MarketData indicates whether to include market data
	MarketData bool `json:"market_data,omitempty"`
	// CommunityData indicates whether to include community data
	CommunityData bool `json:"community_data,omitempty"`
	// DeveloperData indicates whether to include developer data
	DeveloperData bool `json:"developer_data,omitempty"`
	// Sparkline indicates whether to include sparkline data
	Sparkline bool `json:"sparkline,omitempty"`
	// Locale is the language to use for localization
	Locale string `json:"locale,omitempty"`
}

// GetCoinDataByIDResponse represents the response from the Coin Data by ID API
type GetCoinDataByIDResponse struct {
	// ID is the unique identifier of the coin
	ID string `json:"id"`
	// Symbol is the symbol of the coin
	Symbol string `json:"symbol"`
	// Name is the name of the coin
	Name string `json:"name"`
	// Localization contains localized names and descriptions
	Localization map[string]string `json:"localization,omitempty"`
	// MarketData contains market-related data
	MarketData *MarketData `json:"market_data,omitempty"`
	// CommunityData contains community-related data
	CommunityData *CommunityData `json:"community_data,omitempty"`
	// DeveloperData contains developer-related data
	DeveloperData *DeveloperData `json:"developer_data,omitempty"`
}

// MarketData represents market-related data for a coin
type MarketData struct {
	// CurrentPrice contains current prices in different currencies
	CurrentPrice map[string]float64 `json:"current_price"`
	// MarketCap contains market caps in different currencies
	MarketCap map[string]float64 `json:"market_cap"`
	// TotalVolume contains total volumes in different currencies
	TotalVolume map[string]float64 `json:"total_volume"`
	// High24H contains 24h high prices in different currencies
	High24H map[string]float64 `json:"high_24h"`
	// Low24H contains 24h low prices in different currencies
	Low24H map[string]float64 `json:"low_24h"`
	// PriceChange24H contains 24h price changes in different currencies
	PriceChange24H map[string]float64 `json:"price_change_24h"`
	// PriceChangePercentage24H contains 24h price change percentages in different currencies
	PriceChangePercentage24H map[string]float64 `json:"price_change_percentage_24h"`
}

// CommunityData represents community-related data for a coin
type CommunityData struct {
	// FacebookLikes is the number of Facebook likes
	FacebookLikes int `json:"facebook_likes"`
	// TwitterFollowers is the number of Twitter followers
	TwitterFollowers int `json:"twitter_followers"`
	// RedditSubscribers is the number of Reddit subscribers
	RedditSubscribers int `json:"reddit_subscribers"`
	// TelegramChannelUserCount is the number of Telegram channel users
	TelegramChannelUserCount int `json:"telegram_channel_user_count"`
}

// DeveloperData represents developer-related data for a coin
type DeveloperData struct {
	// Forks is the number of forks
	Forks int `json:"forks"`
	// Stars is the number of stars
	Stars int `json:"stars"`
	// Subscribers is the number of subscribers
	Subscribers int `json:"subscribers"`
	// TotalIssues is the total number of issues
	TotalIssues int `json:"total_issues"`
	// ClosedIssues is the number of closed issues
	ClosedIssues int `json:"closed_issues"`
	// PullRequestsMerged is the number of merged pull requests
	PullRequestsMerged int `json:"pull_requests_merged"`
	// PullRequestContributors is the number of pull request contributors
	PullRequestContributors int `json:"pull_request_contributors"`
}

// GetCoinTickersByIDRequest represents the request parameters for getting coin tickers by ID
type GetCoinTickersByIDRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// ExchangeIDs is a list of exchange IDs to filter by
	ExchangeIDs []string `json:"exchange_ids,omitempty"`
	// Include is a list of fields to include in the response
	Include []string `json:"include,omitempty"`
	// Order is the order to sort the results by
	Order string `json:"order,omitempty"`
	// Depth is the depth of the order book to include
	Depth bool `json:"depth,omitempty"`
}

// GetCoinTickersByIDResponse represents the response from the Coin Tickers API
type GetCoinTickersByIDResponse struct {
	Name    string   `json:"name"`
	Tickers []Ticker `json:"tickers"`
}

// Ticker represents a single ticker entry
type Ticker struct {
	Base                   string         `json:"base"`
	Target                 string         `json:"target"`
	Market                 Market         `json:"market"`
	Last                   float64        `json:"last"`
	Volume                 float64        `json:"volume"`
	ConvertedLast          ConvertedPrice `json:"converted_last"`
	ConvertedVolume        ConvertedPrice `json:"converted_volume"`
	TrustScore             string         `json:"trust_score"`
	BidAskSpreadPercentage float64        `json:"bid_ask_spread_percentage"`
	Timestamp              string         `json:"timestamp"`
	LastTradedAt           string         `json:"last_traded_at"`
	LastFetchAt            string         `json:"last_fetch_at"`
	IsAnomaly              bool           `json:"is_anomaly"`
	IsStale                bool           `json:"is_stale"`
	TradeURL               string         `json:"trade_url"`
	TokenInfoURL           string         `json:"token_info_url"`
	CoinID                 string         `json:"coin_id"`
	TargetCoinID           string         `json:"target_coin_id"`
}

// Market represents market information
type Market struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// ConvertedPrice represents converted price information
type ConvertedPrice struct {
	BTC float64 `json:"btc"`
	ETH float64 `json:"eth"`
	USD float64 `json:"usd"`
}

// GetCoinHistoryByIDRequest represents the request parameters for getting coin history by ID
type GetCoinHistoryByIDRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// Date is the date to get history for (dd-mm-yyyy)
	Date string `json:"date" validate:"required"`
	// Localization indicates whether to include localized data
	Localization bool `json:"localization,omitempty"`
}

// GetCoinHistoryByIDResponse represents the response from the Coin History API
type GetCoinHistoryByIDResponse struct {
	ID            string            `json:"id"`
	Symbol        string            `json:"symbol"`
	Name          string            `json:"name"`
	Localization  map[string]string `json:"localization"`
	Image         map[string]string `json:"image"`
	MarketData    MarketData        `json:"market_data"`
	CommunityData CommunityData     `json:"community_data"`
	DeveloperData DeveloperData     `json:"developer_data"`
}

// GetCoinMarketChartByIDRequest represents the request parameters for getting coin market chart by ID
type GetCoinMarketChartByIDRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// VsCurrency is the target currency to get prices in
	VsCurrency string `json:"vs_currency" validate:"required"`
	// Days is the number of days of data to get
	Days string `json:"days" validate:"required"`
	// Interval is the data interval (daily, hourly, minutely)
	Interval string `json:"interval,omitempty"`
}

// GetCoinMarketChartByIDResponse represents the response from the Coin Market Chart API
type GetCoinMarketChartByIDResponse struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

// GetCoinMarketChartRangeRequest represents the request parameters for getting coin market chart by ID within a date range
type GetCoinMarketChartRangeRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// VsCurrency is the target currency to get prices in
	VsCurrency string `json:"vs_currency" validate:"required"`
	// From is the start date (Unix timestamp)
	From int64 `json:"from" validate:"required"`
	// To is the end date (Unix timestamp)
	To int64 `json:"to" validate:"required"`
	// Interval is the data interval (daily, hourly, minutely)
	Interval string `json:"interval,omitempty"`
}

// GetCoinMarketChartRangeResponse represents the response from the Coin Market Chart Range API
type GetCoinMarketChartRangeResponse struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

// GetCoinOHLCByIDRequest represents the request parameters for getting coin OHLC by ID
type GetCoinOHLCByIDRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// VsCurrency is the target currency to get prices in
	VsCurrency string `json:"vs_currency" validate:"required"`
	// Days is the number of days of data to get
	Days string `json:"days" validate:"required"`
	// Interval is the data interval (daily, hourly, minutely)
	Interval string `json:"interval,omitempty"`
}

// GetCoinOHLCByIDResponse represents the response from the Coin OHLC API
type GetCoinOHLCByIDResponse [][]float64

// GetCoinOHLCRangeRequest represents the request parameters for getting coin OHLC by ID within a date range
type GetCoinOHLCRangeRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// VsCurrency is the target currency to get prices in
	VsCurrency string `json:"vs_currency" validate:"required"`
	// From is the start date (Unix timestamp)
	From int64 `json:"from" validate:"required"`
	// To is the end date (Unix timestamp)
	To int64 `json:"to" validate:"required"`
	// Interval is the data interval (daily, hourly, minutely)
	Interval string `json:"interval,omitempty"`
}

// GetCoinOHLCRangeResponse represents the response from the Coin OHLC Range API
type GetCoinOHLCRangeResponse [][]float64

// GetCoinCirculatingSupplyChartRequest represents the request parameters for getting coin circulating supply chart by ID
type GetCoinCirculatingSupplyChartRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// Days is the number of days of data to get
	Days string `json:"days" validate:"required"`
}

// GetCoinCirculatingSupplyChartResponse represents the response from the Coin Circulating Supply Chart API
type GetCoinCirculatingSupplyChartResponse [][]float64

// GetCoinCirculatingSupplyChartRangeRequest represents the request parameters for getting coin circulating supply chart by ID within a date range
type GetCoinCirculatingSupplyChartRangeRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// From is the start date (Unix timestamp)
	From int64 `json:"from" validate:"required"`
	// To is the end date (Unix timestamp)
	To int64 `json:"to" validate:"required"`
}

// GetCoinCirculatingSupplyChartRangeResponse represents the response from the Coin Circulating Supply Chart Range API
type GetCoinCirculatingSupplyChartRangeResponse [][]float64

// GetCoinTotalSupplyChartRequest represents the request parameters for getting coin total supply chart by ID
type GetCoinTotalSupplyChartRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// Days is the number of days of data to get
	Days string `json:"days" validate:"required"`
}

// GetCoinTotalSupplyChartResponse represents the response from the Coin Total Supply Chart API
type GetCoinTotalSupplyChartResponse [][]float64

// GetCoinTotalSupplyChartRangeRequest represents the request parameters for getting coin total supply chart by ID within a date range
type GetCoinTotalSupplyChartRangeRequest struct {
	// ID is the coin ID
	ID string `json:"id" validate:"required"`
	// From is the start date (Unix timestamp)
	From int64 `json:"from" validate:"required"`
	// To is the end date (Unix timestamp)
	To int64 `json:"to" validate:"required"`
}

// GetCoinTotalSupplyChartRangeResponse represents the response from the Coin Total Supply Chart Range API
type GetCoinTotalSupplyChartRangeResponse [][]float64

// Validate validates the request structs
func (r *GetCoinsListRequest) Validate() error {
	// Set default value for Status if empty
	if r.Status == "" {
		r.Status = "active"
	}

	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinsListWithMarketDataRequest) Validate() error {
	// Set default value for Order if empty
	if r.Order == "" {
		r.Order = "market_cap_desc"
	}

	// Set default value for PerPage if empty
	if r.PerPage == 0 {
		r.PerPage = 100
	}

	// Set default value for Page if empty
	if r.Page == 0 {
		r.Page = 1
	}

	// Set default value for Locale if empty
	if r.Locale == "" {
		r.Locale = "en"
	}

	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinDataByIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinTickersByIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinHistoryByIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinMarketChartByIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinMarketChartRangeRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinOHLCByIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinOHLCRangeRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinCirculatingSupplyChartRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinCirculatingSupplyChartRangeRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinTotalSupplyChartRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetCoinTotalSupplyChartRangeRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *GetTopGainersAndLosersRequest) Validate() error {
	// Set default value for Duration if empty
	if r.Duration == "" {
		r.Duration = "24h"
	}

	// Set default value for TopCoins if empty
	if r.TopCoins == 0 {
		r.TopCoins = 1000
	}

	validate := validator.New()
	return validate.Struct(r)
}
