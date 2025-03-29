package nfts

import "github.com/go-playground/validator/v10"

// NFT represents a single NFT
type NFT struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// ContractAddress is the contract address of the NFT
	ContractAddress string `json:"contract_address"`
	// AssetPlatformID is the asset platform ID
	AssetPlatformID string `json:"asset_platform_id"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
}

// GetNFTsListResponse represents the response from the NFTs List API
type GetNFTsListResponse []NFT

// GetNFTDataRequest represents the request parameters for getting NFT data
type GetNFTDataRequest struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id" validate:"required"`
}

// GetNFTDataResponse represents the response from the NFT Data API
type GetNFTDataResponse struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// ContractAddress is the contract address of the NFT
	ContractAddress string `json:"contract_address"`
	// AssetPlatformID is the asset platform ID
	AssetPlatformID string `json:"asset_platform_id"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
	// Image is the image URL of the NFT
	Image string `json:"image"`
	// Description is the description of the NFT
	Description string `json:"description"`
	// ExternalLink is the external link of the NFT
	ExternalLink string `json:"external_link"`
	// Categories is the list of categories
	Categories []string `json:"categories"`
	// FloorPrice is the floor price of the NFT
	FloorPrice map[string]float64 `json:"floor_price"`
	// MarketCap is the market cap of the NFT
	MarketCap map[string]float64 `json:"market_cap"`
	// Volume24H is the 24-hour volume
	Volume24H map[string]float64 `json:"volume_24h"`
	// NumberOfUniqueAddresses is the number of unique addresses
	NumberOfUniqueAddresses int `json:"number_of_unique_addresses"`
	// NumberOfOwners is the number of owners
	NumberOfOwners int `json:"number_of_owners"`
	// TotalSupply is the total supply
	TotalSupply int `json:"total_supply"`
	// CirculatingSupply is the circulating supply
	CirculatingSupply int `json:"circulating_supply"`
	// CreatedAt is the creation timestamp
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the last update timestamp
	UpdatedAt string `json:"updated_at"`
}

// GetNFTContractDataRequest represents the request parameters for getting NFT contract data
type GetNFTContractDataRequest struct {
	// AssetPlatformID is the asset platform ID
	AssetPlatformID string `json:"asset_platform_id" validate:"required"`
	// ContractAddress is the contract address of the NFT
	ContractAddress string `json:"contract_address" validate:"required"`
}

// GetNFTContractDataResponse represents the response from the NFT Contract Data API
type GetNFTContractDataResponse struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// ContractAddress is the contract address of the NFT
	ContractAddress string `json:"contract_address"`
	// AssetPlatformID is the asset platform ID
	AssetPlatformID string `json:"asset_platform_id"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
	// Image is the image URL of the NFT
	Image string `json:"image"`
	// Description is the description of the NFT
	Description string `json:"description"`
	// ExternalLink is the external link of the NFT
	ExternalLink string `json:"external_link"`
	// Categories is the list of categories
	Categories []string `json:"categories"`
	// FloorPrice is the floor price of the NFT
	FloorPrice map[string]float64 `json:"floor_price"`
	// MarketCap is the market cap of the NFT
	MarketCap map[string]float64 `json:"market_cap"`
	// Volume24H is the 24-hour volume
	Volume24H map[string]float64 `json:"volume_24h"`
	// NumberOfUniqueAddresses is the number of unique addresses
	NumberOfUniqueAddresses int `json:"number_of_unique_addresses"`
	// NumberOfOwners is the number of owners
	NumberOfOwners int `json:"number_of_owners"`
	// TotalSupply is the total supply
	TotalSupply int `json:"total_supply"`
	// CirculatingSupply is the circulating supply
	CirculatingSupply int `json:"circulating_supply"`
	// CreatedAt is the creation timestamp
	CreatedAt string `json:"created_at"`
	// UpdatedAt is the last update timestamp
	UpdatedAt string `json:"updated_at"`
}

// GetNFTsMarketDataRequest represents the request parameters for getting NFTs market data
type GetNFTsMarketDataRequest struct {
	// Order is the order to sort by
	Order string `json:"order,omitempty" validate:"omitempty,oneof=market_cap_desc market_cap_asc volume_desc volume_asc id_desc id_asc"`
	// PerPage is the number of results per page
	PerPage int `json:"per_page,omitempty" validate:"omitempty,min=1,max=250"`
	// Page is the page number
	Page int `json:"page,omitempty" validate:"omitempty,min=1"`
	// Sparkline is whether to include sparkline data
	Sparkline bool `json:"sparkline,omitempty"`
}

// GetNFTsMarketDataResponse represents the response from the NFTs Market Data API
type GetNFTsMarketDataResponse []NFTMarketData

// NFTMarketData represents market data for a single NFT
type NFTMarketData struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Image is the image URL of the NFT
	Image string `json:"image"`
	// MarketCap is the market cap of the NFT
	MarketCap map[string]float64 `json:"market_cap"`
	// MarketCapRank is the market cap rank
	MarketCapRank int `json:"market_cap_rank"`
	// FullyDilutedValuation is the fully diluted valuation
	FullyDilutedValuation map[string]float64 `json:"fully_diluted_valuation"`
	// TotalVolume is the total volume
	TotalVolume map[string]float64 `json:"total_volume"`
	// High24H is the 24-hour high
	High24H map[string]float64 `json:"high_24h"`
	// Low24H is the 24-hour low
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
	CirculatingSupply int `json:"circulating_supply"`
	// TotalSupply is the total supply
	TotalSupply int `json:"total_supply"`
	// MaxSupply is the maximum supply
	MaxSupply int `json:"max_supply"`
	// Ath is the all-time high
	Ath map[string]float64 `json:"ath"`
	// AthChangePercentage is the all-time high change percentage
	AthChangePercentage map[string]float64 `json:"ath_change_percentage"`
	// AthDate is the all-time high date
	AthDate map[string]string `json:"ath_date"`
	// Atl is the all-time low
	Atl map[string]float64 `json:"atl"`
	// AtlChangePercentage is the all-time low change percentage
	AtlChangePercentage map[string]float64 `json:"atl_change_percentage"`
	// AtlDate is the all-time low date
	AtlDate map[string]string `json:"atl_date"`
	// LastUpdated is the last update timestamp
	LastUpdated string `json:"last_updated"`
	// SparklineIn7D is the 7-day sparkline data
	SparklineIn7D *Sparkline `json:"sparkline_in_7d,omitempty"`
}

// GetNFTHistoryRequest represents the request parameters for getting NFT history
type GetNFTHistoryRequest struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id" validate:"required"`
	// VsCurrency is the target currency of market data
	VsCurrency string `json:"vs_currency" validate:"required"`
	// Days is the data up to number of days ago
	Days string `json:"days" validate:"required"`
}

// GetNFTHistoryResponse represents the response from the NFT History API
type GetNFTHistoryResponse struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Image is the image URL of the NFT
	Image string `json:"image"`
	// MarketData is the market data
	MarketData *NFTMarketData `json:"market_data"`
	// CommunityData is the community data
	CommunityData *NFTCommunityData `json:"community_data"`
	// DeveloperData is the developer data
	DeveloperData *NFTDeveloperData `json:"developer_data"`
	// PublicInterestStats is the public interest stats
	PublicInterestStats *NFTPublicInterestStats `json:"public_interest_stats"`
}

// GetNFTContractHistoryRequest represents the request parameters for getting NFT contract history
type GetNFTContractHistoryRequest struct {
	// AssetPlatformID is the asset platform ID
	AssetPlatformID string `json:"asset_platform_id" validate:"required"`
	// ContractAddress is the contract address of the NFT
	ContractAddress string `json:"contract_address" validate:"required"`
	// VsCurrency is the target currency of market data
	VsCurrency string `json:"vs_currency" validate:"required"`
	// Days is the data up to number of days ago
	Days string `json:"days" validate:"required"`
}

// GetNFTContractHistoryResponse represents the response from the NFT Contract History API
type GetNFTContractHistoryResponse struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Image is the image URL of the NFT
	Image string `json:"image"`
	// MarketData is the market data
	MarketData *NFTMarketData `json:"market_data"`
	// CommunityData is the community data
	CommunityData *NFTCommunityData `json:"community_data"`
	// DeveloperData is the developer data
	DeveloperData *NFTDeveloperData `json:"developer_data"`
	// PublicInterestStats is the public interest stats
	PublicInterestStats *NFTPublicInterestStats `json:"public_interest_stats"`
}

// GetNFTTickersRequest represents the request parameters for getting NFT tickers
type GetNFTTickersRequest struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id" validate:"required"`
	// ExchangeIDs is the list of exchange IDs to filter by
	ExchangeIDs string `json:"exchange_ids,omitempty"`
	// IncludeExchangeLogo is whether to include exchange logo
	IncludeExchangeLogo bool `json:"include_exchange_logo,omitempty"`
	// Order is the order to sort by
	Order string `json:"order,omitempty" validate:"omitempty,oneof=trust_score_desc trust_score_asc volume_desc volume_asc"`
	// Depth is the depth of order book
	Depth bool `json:"depth,omitempty"`
}

// GetNFTTickersResponse represents the response from the NFT Tickers API
type GetNFTTickersResponse struct {
	// Name is the name of the NFT
	Name string `json:"name"`
	// Tickers is the list of tickers
	Tickers []Ticker `json:"tickers"`
}

// Ticker represents a single ticker
type Ticker struct {
	// Base is the base currency
	Base string `json:"base"`
	// Target is the target currency
	Target string `json:"target"`
	// Market is the market information
	Market *Market `json:"market"`
	// Last is the last price
	Last float64 `json:"last"`
	// Volume is the volume
	Volume float64 `json:"volume"`
	// ConvertedLast is the converted last price
	ConvertedLast map[string]float64 `json:"converted_last"`
	// ConvertedVolume is the converted volume
	ConvertedVolume map[string]float64 `json:"converted_volume"`
	// TrustScore is the trust score
	TrustScore string `json:"trust_score"`
	// BidAskSpreadPercentage is the bid-ask spread percentage
	BidAskSpreadPercentage float64 `json:"bid_ask_spread_percentage"`
	// Timestamp is the timestamp
	Timestamp string `json:"timestamp"`
	// LastTradedAt is the last traded timestamp
	LastTradedAt string `json:"last_traded_at"`
	// LastFetchAt is the last fetch timestamp
	LastFetchAt string `json:"last_fetch_at"`
	// IsAnomaly is whether the ticker is anomalous
	IsAnomaly bool `json:"is_anomaly"`
	// IsStale is whether the ticker is stale
	IsStale bool `json:"is_stale"`
	// TradeURL is the trade URL
	TradeURL string `json:"trade_url"`
	// TokenInfoURL is the token info URL
	TokenInfoURL string `json:"token_info_url"`
	// CoinID is the coin ID
	CoinID string `json:"coin_id"`
	// TargetCoinID is the target coin ID
	TargetCoinID string `json:"target_coin_id"`
}

// Market represents market information
type Market struct {
	// Identifier is the market identifier
	Identifier string `json:"identifier"`
	// Name is the market name
	Name string `json:"name"`
	// HasTradingIncentive is whether the market has trading incentive
	HasTradingIncentive bool `json:"has_trading_incentive"`
}

// NFTCommunityData represents community data for a single NFT
type NFTCommunityData struct {
	// TwitterFollowers is the number of Twitter followers
	TwitterFollowers int `json:"twitter_followers"`
	// RedditSubscribers is the number of Reddit subscribers
	RedditSubscribers int `json:"reddit_subscribers"`
	// TelegramChannelUserCount is the number of Telegram channel users
	TelegramChannelUserCount int `json:"telegram_channel_user_count"`
}

// NFTDeveloperData represents developer data for a single NFT
type NFTDeveloperData struct {
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
	// CodeAdditionsDeletions4Weeks is the code additions and deletions in the last 4 weeks
	CodeAdditionsDeletions4Weeks *CodeAdditionsDeletions `json:"code_additions_deletions_4_weeks"`
	// CommitCount4Weeks is the number of commits in the last 4 weeks
	CommitCount4Weeks int `json:"commit_count_4_weeks"`
}

// NFTPublicInterestStats represents public interest stats for a single NFT
type NFTPublicInterestStats struct {
	// AlexaRank is the Alexa rank
	AlexaRank int `json:"alexa_rank"`
	// BingMatches is the number of Bing matches
	BingMatches int `json:"bing_matches"`
}

// CodeAdditionsDeletions represents code additions and deletions
type CodeAdditionsDeletions struct {
	// Additions is the number of additions
	Additions []int `json:"additions"`
	// Deletions is the number of deletions
	Deletions []int `json:"deletions"`
}

// Sparkline represents sparkline data
type Sparkline struct {
	// Price is the list of prices
	Price []float64 `json:"price"`
}

// Validate validates the request parameters
func (r *GetNFTDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetNFTContractDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetNFTsMarketDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetNFTHistoryRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetNFTContractHistoryRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Validate validates the request parameters
func (r *GetNFTTickersRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
