# CoinGecko V3 API SDK

A Go language SDK for accessing the CoinGecko V3 API. This SDK provides a simple, type-safe way to interact with the CoinGecko API.

[中文文档](README_CN.md)

## Features

- Complete CoinGecko V3 API support
- Type-safe request and response handling
- Unified error handling
- API key authentication support
- Automatic retry mechanism
- Request timeout control

## Supported API Endpoints

- `/ping` - API availability check
- `/key` - API key status query
- `/coins` - Cryptocurrency data
- `/asset_platforms` - Asset platform information
- `/categories` - Cryptocurrency categories
- `/exchanges` - Exchange information
- `/contract` - Contract data
- `/derivatives` - Derivatives data
- `/nfts` - NFT-related data
- `/exchange_rates` - Exchange rate information
- `/search` - Search functionality
- `/trending` - Trending data
- `/global` - Global market data
- `/companies` - Company data

## Development Status

⚠️ **Note: This project is currently in development and testing phase**

- Basic functionality has been implemented
- Comprehensive testing is in progress
- API interfaces may change
- Recommended for use in testing environments only

## Installation

```bash
go get github.com/ipangpang/coingecko-v3
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/ipangpang/coingecko-v3/pkg"
)

func main() {
    // Create client
    client := pkg.NewClient()

    // Check API availability
    ping, err := client.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("API Status: %s\n", ping.GeckoSays)

    // Get Bitcoin price
    price, err := client.GetCoinPriceByIDs(&simple.GetCoinPriceByIDsRequest{
        IDs: []string{"bitcoin"},
        VsCurrencies: []string{"usd"},
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Bitcoin Price: $%v\n", price.Bitcoin.USD)
}
```

## Configuration Options

```go
client := pkg.NewClient(
    pkg.WithBaseURL("https://api.coingecko.com/api/v3"),
    pkg.WithTimeout(10 * time.Second),
    pkg.WithAPIKey("your-api-key"),
)
```

## Contributing

Issues and Pull Requests are welcome to help improve this project.

## License

MIT License 