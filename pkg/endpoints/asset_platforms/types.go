package asset_platforms

import "github.com/go-playground/validator/v10"

// GetAssetPlatformsResponse represents the response from the Asset Platforms List API
type GetAssetPlatformsResponse []AssetPlatform

// AssetPlatform represents a single asset platform
type AssetPlatform struct {
	// ID is the unique identifier of the asset platform
	ID string `json:"id"`
	// ChainIdentifier is the chain identifier
	ChainIdentifier int `json:"chain_identifier,omitempty"`
	// Name is the name of the asset platform
	Name string `json:"name"`
	// ShortName is the short name of the asset platform
	ShortName string `json:"short_name,omitempty"`
}

// GetTokenListsByAssetPlatformIDRequest represents the request parameters for getting token lists by asset platform ID
type GetTokenListsByAssetPlatformIDRequest struct {
	// AssetPlatformID is the ID of the asset platform
	AssetPlatformID string `json:"asset_platform_id" validate:"required"`
}

// GetTokenListsByAssetPlatformIDResponse represents the response from the Token Lists by Asset Platform ID API
type GetTokenListsByAssetPlatformIDResponse []Token

// Token represents a single token
type Token struct {
	// ID is the unique identifier of the token
	ID string `json:"id"`
	// Symbol is the symbol of the token
	Symbol string `json:"symbol"`
	// Name is the name of the token
	Name string `json:"name"`
	// ContractAddress is the contract address of the token
	ContractAddress string `json:"contract_address"`
}

// Validate validates the request parameters
func (r *GetTokenListsByAssetPlatformIDRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
