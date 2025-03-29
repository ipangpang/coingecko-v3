package pkg

import (
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/asset_platforms"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/categories"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/coins"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/companies"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/contract"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/derivatives"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/exchange_rates"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/exchanges"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/global"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/key"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/nfts"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/ping"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/search"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/simple"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/trending"
)

// ClientOption defines client configuration options
type ClientOption = base.ClientOption

type Client interface {
	Ping() (*ping.PingResponse, error)
	Key() (*key.KeyResponse, error)
	GetCoinsList(request *coins.GetCoinsListRequest) (*coins.GetCoinsListResponse, error)
	GetCoinPriceByIDs(request *simple.GetCoinPriceByIDsRequest) (*simple.GetCoinPriceByIDsResponse, error)
	GetCoinPriceByTokenAddress(request *simple.GetCoinPriceByTokenAddressRequest) (*simple.GetCoinPriceByTokenAddressResponse, error)
	GetSupportedCurrencies() (*simple.GetSupportedCurrenciesResponse, error)
	GetAssetPlatforms() (*asset_platforms.GetAssetPlatformsResponse, error)
	GetCategoriesList() (*categories.GetCategoriesListResponse, error)
	GetCategoriesData(request *categories.GetCategoriesDataRequest) (*categories.GetCategoriesDataResponse, error)
	GetExchangesList(request *exchanges.GetExchangesListRequest) (*exchanges.GetExchangesListResponse, error)
	GetExchangeData(request *exchanges.GetExchangeDataRequest) (*exchanges.GetExchangeDataResponse, error)
	GetExchangeTickers(request *exchanges.GetExchangeTickersRequest) (*exchanges.GetExchangeTickersResponse, error)
	GetExchangeVolumeChart(request *exchanges.GetExchangeVolumeChartRequest) (*exchanges.GetExchangeVolumeChartResponse, error)
	GetContractData(request *contract.GetContractDataRequest) (*contract.GetContractDataResponse, error)
	GetContractMarketChart(request *contract.GetContractMarketChartRequest) (*contract.GetContractMarketChartResponse, error)
	GetContractMarketChartRange(request *contract.GetContractMarketChartRangeRequest) (*contract.GetContractMarketChartRangeResponse, error)
	GetDerivativesList() (*derivatives.GetDerivativesListResponse, error)
	GetDerivativesExchangesList() (*derivatives.GetDerivativesExchangesListResponse, error)
	GetDerivativeExchangeData(request *derivatives.GetDerivativeExchangeDataRequest) (*derivatives.GetDerivativeExchangeDataResponse, error)
	GetDerivativesExchangesListIDMap() (*derivatives.GetDerivativesExchangesListIDMapResponse, error)
	GetNFTsList() (*nfts.GetNFTsListResponse, error)
	GetNFTData(request *nfts.GetNFTDataRequest) (*nfts.GetNFTDataResponse, error)
	GetNFTContractData(request *nfts.GetNFTContractDataRequest) (*nfts.GetNFTContractDataResponse, error)
	GetNFTsMarketData(request *nfts.GetNFTsMarketDataRequest) (*nfts.GetNFTsMarketDataResponse, error)
	GetNFTHistory(request *nfts.GetNFTHistoryRequest) (*nfts.GetNFTHistoryResponse, error)
	GetNFTContractHistory(request *nfts.GetNFTContractHistoryRequest) (*nfts.GetNFTContractHistoryResponse, error)
	GetNFTTickers(request *nfts.GetNFTTickersRequest) (*nfts.GetNFTTickersResponse, error)
	GetExchangeRates() (*exchange_rates.GetExchangeRatesResponse, error)
	Search(request *search.SearchRequest) (*search.SearchResponse, error)
	GetTrending() (*trending.TrendingResponse, error)
	GetGlobal() (*global.GetGlobalResponse, error)
	GetGlobalDefi() (*global.GetGlobalDefiResponse, error)
	GetGlobalMarketCapChart(vsCurrency string, days string) (*global.GetGlobalMarketCapChartResponse, error)
	GetPublicTreasury() (*companies.GetPublicTreasuryResponse, error)
}

type ClientImpl struct {
	PingClient           ping.Client
	KeyClient            key.Client
	CoinsClient          coins.Client
	SampleClient         simple.Client
	AssetPlatformsClient asset_platforms.Client
	CategoriesClient     categories.Client
	ExchangesClient      exchanges.Client
	ContractClient       contract.Client
	DerivativesClient    derivatives.Client
	NFTsClient           nfts.Client
	ExchangeRatesClient  exchange_rates.Client
	SearchClient         search.Client
	TrendingClient       trending.Client
	GlobalClient         global.Client
	CompaniesClient      companies.Client
}

func NewClient(options ...ClientOption) Client {
	config := base.DefaultConfig()
	baseClient := base.NewBaseClient(config, options...)

	client := &ClientImpl{}
	client.PingClient = ping.NewClient(baseClient)
	client.KeyClient = key.NewClient(baseClient)
	client.CoinsClient = coins.NewClient(baseClient)
	client.SampleClient = simple.NewClient(baseClient)
	client.AssetPlatformsClient = asset_platforms.NewClient(baseClient)
	client.CategoriesClient = categories.NewClient(baseClient)
	client.ExchangesClient = exchanges.NewClient(baseClient)
	client.ContractClient = contract.NewClient(baseClient)
	client.DerivativesClient = derivatives.NewClient(baseClient)
	client.NFTsClient = nfts.NewClient(baseClient)
	client.ExchangeRatesClient = exchange_rates.NewClient(baseClient)
	client.SearchClient = search.NewClient(baseClient)
	client.TrendingClient = trending.NewClient(baseClient)
	client.GlobalClient = global.NewClient(baseClient)
	client.CompaniesClient = companies.NewClient(baseClient)

	return client
}

func (c ClientImpl) Ping() (*ping.PingResponse, error) {
	return c.PingClient.Ping()
}

func (c ClientImpl) Key() (*key.KeyResponse, error) {
	return c.KeyClient.Key()
}

func (c ClientImpl) GetCoinsList(request *coins.GetCoinsListRequest) (*coins.GetCoinsListResponse, error) {
	return c.CoinsClient.GetCoinsList(request)
}

func (c ClientImpl) GetCoinPriceByIDs(request *simple.GetCoinPriceByIDsRequest) (*simple.GetCoinPriceByIDsResponse, error) {
	return c.SampleClient.GetCoinPriceByIDs(request)
}

func (c ClientImpl) GetCoinPriceByTokenAddress(request *simple.GetCoinPriceByTokenAddressRequest) (*simple.GetCoinPriceByTokenAddressResponse, error) {
	return c.SampleClient.GetCoinPriceByTokenAddress(request)
}

func (c ClientImpl) GetSupportedCurrencies() (*simple.GetSupportedCurrenciesResponse, error) {
	return c.SampleClient.GetSupportedCurrencies()
}

func (c ClientImpl) GetAssetPlatforms() (*asset_platforms.GetAssetPlatformsResponse, error) {
	return c.AssetPlatformsClient.GetAssetPlatforms()
}

func (c ClientImpl) GetCategoriesList() (*categories.GetCategoriesListResponse, error) {
	return c.CategoriesClient.GetCategoriesList()
}

func (c ClientImpl) GetCategoriesData(request *categories.GetCategoriesDataRequest) (*categories.GetCategoriesDataResponse, error) {
	return c.CategoriesClient.GetCategoriesData(request)
}

func (c ClientImpl) GetExchangesList(request *exchanges.GetExchangesListRequest) (*exchanges.GetExchangesListResponse, error) {
	return c.ExchangesClient.GetExchangesList(request)
}

func (c ClientImpl) GetExchangeData(request *exchanges.GetExchangeDataRequest) (*exchanges.GetExchangeDataResponse, error) {
	return c.ExchangesClient.GetExchangeData(request)
}

func (c ClientImpl) GetExchangeTickers(request *exchanges.GetExchangeTickersRequest) (*exchanges.GetExchangeTickersResponse, error) {
	return c.ExchangesClient.GetExchangeTickers(request)
}

func (c ClientImpl) GetExchangeVolumeChart(request *exchanges.GetExchangeVolumeChartRequest) (*exchanges.GetExchangeVolumeChartResponse, error) {
	return c.ExchangesClient.GetExchangeVolumeChart(request)
}

func (c ClientImpl) GetContractData(request *contract.GetContractDataRequest) (*contract.GetContractDataResponse, error) {
	return c.ContractClient.GetContractData(request)
}

func (c ClientImpl) GetContractMarketChart(request *contract.GetContractMarketChartRequest) (*contract.GetContractMarketChartResponse, error) {
	return c.ContractClient.GetContractMarketChart(request)
}

func (c ClientImpl) GetContractMarketChartRange(request *contract.GetContractMarketChartRangeRequest) (*contract.GetContractMarketChartRangeResponse, error) {
	return c.ContractClient.GetContractMarketChartRange(request)
}

func (c ClientImpl) GetDerivativesList() (*derivatives.GetDerivativesListResponse, error) {
	return c.DerivativesClient.GetDerivativesList()
}

func (c ClientImpl) GetDerivativesExchangesList() (*derivatives.GetDerivativesExchangesListResponse, error) {
	return c.DerivativesClient.GetDerivativesExchangesList()
}

func (c ClientImpl) GetDerivativeExchangeData(request *derivatives.GetDerivativeExchangeDataRequest) (*derivatives.GetDerivativeExchangeDataResponse, error) {
	return c.DerivativesClient.GetDerivativeExchangeData(request)
}

func (c ClientImpl) GetDerivativesExchangesListIDMap() (*derivatives.GetDerivativesExchangesListIDMapResponse, error) {
	return c.DerivativesClient.GetDerivativesExchangesListIDMap()
}

func (c ClientImpl) GetNFTsList() (*nfts.GetNFTsListResponse, error) {
	return c.NFTsClient.GetNFTsList()
}

func (c ClientImpl) GetNFTData(request *nfts.GetNFTDataRequest) (*nfts.GetNFTDataResponse, error) {
	return c.NFTsClient.GetNFTData(request)
}

func (c ClientImpl) GetNFTContractData(request *nfts.GetNFTContractDataRequest) (*nfts.GetNFTContractDataResponse, error) {
	return c.NFTsClient.GetNFTContractData(request)
}

func (c ClientImpl) GetNFTsMarketData(request *nfts.GetNFTsMarketDataRequest) (*nfts.GetNFTsMarketDataResponse, error) {
	return c.NFTsClient.GetNFTsMarketData(request)
}

func (c ClientImpl) GetNFTHistory(request *nfts.GetNFTHistoryRequest) (*nfts.GetNFTHistoryResponse, error) {
	return c.NFTsClient.GetNFTHistory(request)
}

func (c ClientImpl) GetNFTContractHistory(request *nfts.GetNFTContractHistoryRequest) (*nfts.GetNFTContractHistoryResponse, error) {
	return c.NFTsClient.GetNFTContractHistory(request)
}

func (c ClientImpl) GetNFTTickers(request *nfts.GetNFTTickersRequest) (*nfts.GetNFTTickersResponse, error) {
	return c.NFTsClient.GetNFTTickers(request)
}

func (c ClientImpl) GetExchangeRates() (*exchange_rates.GetExchangeRatesResponse, error) {
	return c.ExchangeRatesClient.GetExchangeRates()
}

func (c ClientImpl) Search(request *search.SearchRequest) (*search.SearchResponse, error) {
	return c.SearchClient.Search(request)
}

func (c ClientImpl) GetTrending() (*trending.TrendingResponse, error) {
	return c.TrendingClient.GetTrending()
}

func (c ClientImpl) GetGlobal() (*global.GetGlobalResponse, error) {
	return c.GlobalClient.GetGlobal()
}

func (c ClientImpl) GetGlobalDefi() (*global.GetGlobalDefiResponse, error) {
	return c.GlobalClient.GetGlobalDefi()
}

func (c ClientImpl) GetGlobalMarketCapChart(vsCurrency string, days string) (*global.GetGlobalMarketCapChartResponse, error) {
	return c.GlobalClient.GetGlobalMarketCapChart(vsCurrency, days)
}

func (c ClientImpl) GetPublicTreasury() (*companies.GetPublicTreasuryResponse, error) {
	return c.CompaniesClient.GetPublicTreasury()
}
