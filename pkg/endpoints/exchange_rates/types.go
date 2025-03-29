package exchange_rates

import (
	"github.com/go-playground/validator/v10"
)

// ExchangeRate represents a single exchange rate
type ExchangeRate struct {
	// Name is the name of the currency
	Name string `json:"name"`
	// Unit is the unit of the currency
	Unit string `json:"unit"`
	// Value is the value of the currency in BTC
	Value float64 `json:"value"`
	// Type is the type of the currency
	Type string `json:"type"`
}

// GetExchangeRatesResponse represents the response from the Exchange Rates API
type GetExchangeRatesResponse struct {
	// Rates contains exchange rates for different currencies
	Rates map[string]ExchangeRate `json:"rates"`
}

// GetExchangeRateRequest represents the request parameters for getting a specific exchange rate
type GetExchangeRateRequest struct {
	// ID is the unique identifier of the currency
	ID string `json:"id" validate:"required"`
}

// GetExchangeRateResponse represents the response from the Exchange Rate API
type GetExchangeRateResponse struct {
	// Name is the name of the currency
	Name string `json:"name"`
	// Unit is the unit of the currency
	Unit string `json:"unit"`
	// Value is the exchange rate value
	Value float64 `json:"value"`
	// Type is the type of the currency
	Type string `json:"type"`
}

// Validate validates the GetExchangeRateRequest
func (r *GetExchangeRateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
