# CoinGecko V3 API SDK

这是一个用于访问 CoinGecko V3 API 的 Go 语言 SDK。该 SDK 提供了简单、类型安全的方式来与 CoinGecko API 进行交互。

[English Documentation](README.md)

## 功能特点

- 完整的 CoinGecko V3 API 支持
- 类型安全的请求和响应处理
- 统一的错误处理
- 支持 API 密钥认证
- 自动重试机制
- 请求超时控制

## 支持的 API 端点

- `/ping` - API 可用性检查
- `/key` - API 密钥状态查询
- `/coins` - 加密货币相关数据
- `/asset_platforms` - 资产平台信息
- `/categories` - 加密货币分类
- `/exchanges` - 交易所信息
- `/contract` - 合约数据
- `/derivatives` - 衍生品数据
- `/nfts` - NFT 相关数据
- `/exchange_rates` - 汇率信息
- `/search` - 搜索功能
- `/trending` - 趋势数据
- `/global` - 全局市场数据
- `/companies` - 公司数据

## 开发状态

⚠️ **注意：本项目目前处于开发测试阶段**

- 基础功能已经实现
- 正在进行全面测试
- API 接口可能会发生变化
- 建议在测试环境中使用

## 安装

```bash
go get github.com/ipangpang/coingecko-v3
```

## 快速开始

```go
package main

import (
    "fmt"
    "log"
    "github.com/ipangpang/coingecko-v3/pkg"
)

func main() {
    // 创建客户端
    client := pkg.NewClient()

    // 检查 API 可用性
    ping, err := client.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("API Status: %s\n", ping.GeckoSays)

    // 获取比特币价格
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

## 配置选项

```go
client := pkg.NewClient(
    pkg.WithBaseURL("https://api.coingecko.com/api/v3"),
    pkg.WithTimeout(10 * time.Second),
    pkg.WithAPIKey("your-api-key"),
)
```

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进这个项目。

## 许可证

MIT License 