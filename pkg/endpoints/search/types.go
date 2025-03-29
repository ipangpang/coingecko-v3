package search

import "github.com/go-playground/validator/v10"

// SearchRequest represents the request parameters for search
type SearchRequest struct {
	// Query is the search query
	Query string `json:"query" validate:"required"`
}

// SearchResponse represents the response from the Search API
type SearchResponse struct {
	// Coins contains the list of coins
	Coins []Coin `json:"coins"`
	// Exchanges contains the list of exchanges
	Exchanges []Exchange `json:"exchanges"`
	// Categories contains the list of categories
	Categories []Category `json:"categories"`
	// NFTs contains the list of NFTs
	NFTs []NFT `json:"nfts"`
}

// Coin represents a single coin in search results
type Coin struct {
	// ID is the unique identifier of the coin
	ID string `json:"id"`
	// Name is the name of the coin
	Name string `json:"name"`
	// Symbol is the symbol of the coin
	Symbol string `json:"symbol"`
	// MarketCapRank is the rank of the coin by market cap
	MarketCapRank int `json:"market_cap_rank"`
	// Thumb is the thumbnail URL
	Thumb string `json:"thumb"`
	// Large is the large image URL
	Large string `json:"large"`
}

// Exchange represents a single exchange in search results
type Exchange struct {
	// ID is the unique identifier of the exchange
	ID string `json:"id"`
	// Name is the name of the exchange
	Name string `json:"name"`
	// MarketType is the market type of the exchange
	MarketType string `json:"market_type"`
	// Thumb is the thumbnail URL
	Thumb string `json:"thumb"`
	// Large is the large image URL
	Large string `json:"large"`
}

// Category represents a single category in search results
type Category struct {
	// ID is the unique identifier of the category
	ID string `json:"id"`
	// Name is the name of the category
	Name string `json:"name"`
}

// NFT represents a single NFT in search results
type NFT struct {
	// ID is the unique identifier of the NFT
	ID string `json:"id"`
	// Name is the name of the NFT
	Name string `json:"name"`
	// Symbol is the symbol of the NFT
	Symbol string `json:"symbol"`
	// Thumb is the thumbnail URL
	Thumb string `json:"thumb"`
}

// Validate validates the request structs
func (r *SearchRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
