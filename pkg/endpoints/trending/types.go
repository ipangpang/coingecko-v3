package trending

// TrendingResponse represents the response from the Trending API
type TrendingResponse struct {
	// Coins contains the list of trending coins
	Coins []TrendingCoin `json:"coins"`
	// NFTs contains the list of trending NFTs
	NFTs []TrendingNFT `json:"nfts"`
	// Categories contains the list of trending categories
	Categories []TrendingCategory `json:"categories"`
}

// TrendingCoin represents a single trending coin
type TrendingCoin struct {
	// Item contains the coin information
	Item TrendingCoinItem `json:"item"`
}

// TrendingCoinItem represents the details of a trending coin
type TrendingCoinItem struct {
	// ID is the unique identifier of the coin
	ID string `json:"id"`
	// CoinID is the unique identifier of the coin
	CoinID int `json:"coin_id"`
	// Name is the name of the coin
	Name string `json:"name"`
	// Symbol is the symbol of the coin
	Symbol string `json:"symbol"`
	// MarketCapRank is the rank of the coin by market cap
	MarketCapRank int `json:"market_cap_rank"`
	// Thumb is the URL of the coin's thumbnail image
	Thumb string `json:"thumb"`
	// Small is the URL of the coin's small image
	Small string `json:"small"`
	// Large is the URL of the coin's large image
	Large string `json:"large"`
	// Slug is the slug of the coin
	Slug string `json:"slug"`
	// PriceBTC is the price of the coin in BTC
	PriceBTC float64 `json:"price_btc"`
	// Score is the score of the coin
	Score int `json:"score"`
}

// TrendingNFT represents a single trending NFT
type TrendingNFT struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
	// Thumb is the thumbnail URL
	Thumb string `json:"thumb"`
	// NFTContractID is the NFT contract ID
	NFTContractID int `json:"nft_contract_id"`
}

// TrendingCategory represents a single trending category
type TrendingCategory struct {
	// ID is the unique identifier of the category
	ID string `json:"id"`
	// Name is the name of the category
	Name string `json:"name"`
	// MarketCap is the market cap of the category
	MarketCap float64 `json:"market_cap"`
	// MarketCapChange24H is the 24h market cap change
	MarketCapChange24H float64 `json:"market_cap_change_24h"`
	// Content is the content of the category
	Content string `json:"content"`
	// Top3Coins is the list of top 3 coins in the category
	Top3Coins []string `json:"top_3_coins"`
	// Volume24H is the 24h volume
	Volume24H float64 `json:"volume_24h"`
	// UpdatedAt is the last update timestamp
	UpdatedAt string `json:"updated_at"`
}
