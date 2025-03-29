package categories

import "github.com/go-playground/validator/v10"

// Category represents a single category
type Category struct {
	// CategoryID is the unique identifier of the category
	CategoryID string `json:"category_id"`
	// Name is the name of the category
	Name string `json:"name"`
	// MarketCap is the market cap of the category
	MarketCap map[string]float64 `json:"market_cap"`
	// MarketCapChange24H is the 24-hour market cap change
	MarketCapChange24H map[string]float64 `json:"market_cap_change_24h"`
	// Content is the content of the category
	Content string `json:"content,omitempty"`
	// Top3Coins is the list of top 3 coins in the category
	Top3Coins []string `json:"top_3_coins,omitempty"`
	// Volume24H is the 24-hour volume
	Volume24H map[string]float64 `json:"volume_24h"`
	// UpdatedAt is the last update timestamp
	UpdatedAt string `json:"updated_at"`
	// Coins is the list of coins in the category
	Coins []Coin `json:"coins"`
}

// GetCategoriesListResponse represents the response from the Categories List API
type GetCategoriesListResponse []Category

// GetCategoriesDataRequest represents the request parameters for getting categories data
type GetCategoriesDataRequest struct {
	// Order is the order to sort by
	Order string `json:"order,omitempty" validate:"omitempty,oneof=market_cap_desc market_cap_asc name_desc name_asc market_cap_change_24h_desc market_cap_change_24h_asc"`
	// PerPage is the number of results per page
	PerPage int `json:"per_page,omitempty" validate:"omitempty,min=1,max=250"`
	// Page is the page number
	Page int `json:"page,omitempty" validate:"omitempty,min=1"`
	// Sparkline is whether to include sparkline data
	Sparkline bool `json:"sparkline,omitempty"`
	// PriceChangePercentage is the price change percentage to filter by
	PriceChangePercentage string `json:"price_change_percentage,omitempty" validate:"omitempty,oneof=1h 24h 7d 14d 30d 200d 1y"`
}

// GetCategoriesDataResponse represents the response from the Categories Data API
type GetCategoriesDataResponse []Category

// GetCategoryDataRequest represents the request parameters for getting category data
type GetCategoryDataRequest struct {
	// CategoryID is the unique identifier of the category
	CategoryID string `json:"category_id" validate:"required"`
	// Order is the order to sort by
	Order string `json:"order,omitempty" validate:"omitempty,oneof=market_cap_desc market_cap_asc name_desc name_asc market_cap_change_24h_desc market_cap_change_24h_asc"`
	// PerPage is the number of results per page
	PerPage int `json:"per_page,omitempty" validate:"omitempty,min=1,max=250"`
	// Page is the page number
	Page int `json:"page,omitempty" validate:"omitempty,min=1"`
	// Sparkline is whether to include sparkline data
	Sparkline bool `json:"sparkline,omitempty"`
	// PriceChangePercentage is the price change percentage to filter by
	PriceChangePercentage string `json:"price_change_percentage,omitempty" validate:"omitempty,oneof=1h 24h 7d 14d 30d 200d 1y"`
}

// GetCategoryDataResponse represents the response from the Category Data API
type GetCategoryDataResponse struct {
	// CategoryID is the unique identifier of the category
	CategoryID string `json:"category_id"`
	// Name is the name of the category
	Name string `json:"name"`
	// MarketCap is the market cap of the category
	MarketCap map[string]float64 `json:"market_cap"`
	// MarketCapChange24H is the 24-hour market cap change
	MarketCapChange24H map[string]float64 `json:"market_cap_change_24h"`
	// Content is the content of the category
	Content string `json:"content,omitempty"`
	// Top3Coins is the list of top 3 coins in the category
	Top3Coins []string `json:"top_3_coins,omitempty"`
	// Volume24H is the 24-hour volume
	Volume24H map[string]float64 `json:"volume_24h"`
	// UpdatedAt is the last update timestamp
	UpdatedAt string `json:"updated_at"`
	// Coins is the list of coins in the category
	Coins []Coin `json:"coins"`
}

// Coin represents a single coin in a category
type Coin struct {
	// ID is the unique identifier of the coin
	ID string `json:"id"`
	// CoinID is the unique identifier of the coin
	CoinID string `json:"coin_id"`
	// Name is the name of the coin
	Name string `json:"name"`
	// Symbol is the symbol of the coin
	Symbol string `json:"symbol"`
	// MarketCapRank is the rank of the coin by market cap
	MarketCapRank int `json:"market_cap_rank"`
	// Thumb is the thumbnail URL of the coin
	Thumb string `json:"thumb"`
	// Small is the small image URL of the coin
	Small string `json:"small"`
	// Large is the large image URL of the coin
	Large string `json:"large"`
	// Slug is the slug of the coin
	Slug string `json:"slug"`
	// PriceChangePercentage24H is the 24-hour price change percentage
	PriceChangePercentage24H float64 `json:"price_change_percentage_24h"`
	// Sparkline is the sparkline data
	Sparkline *Sparkline `json:"sparkline,omitempty"`
}

// Sparkline represents sparkline data
type Sparkline struct {
	// Price is the list of prices
	Price []float64 `json:"price"`
}

// Validate validates the request parameters
func (r *GetCategoriesDataRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
