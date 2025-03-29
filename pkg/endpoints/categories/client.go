package categories

import (
	"fmt"
	"github.com/ipangpang/coingecko-v3/pkg/endpoints/base"
)

const (
	GetCategoriesListRequestPoint = "/coins/categories/list"
	GetCategoriesDataRequestPoint = "/coins/categories"
)

type Client interface {
	GetCategoriesList() (*GetCategoriesListResponse, error)
	GetCategoriesData(request *GetCategoriesDataRequest) (*GetCategoriesDataResponse, error)
}

type ClientImpl struct {
	baseClient *base.BaseClient
}

func NewClient(baseClient *base.BaseClient) Client {
	return &ClientImpl{
		baseClient: baseClient,
	}
}

func (c *ClientImpl) GetCategoriesList() (*GetCategoriesListResponse, error) {
	var response GetCategoriesListResponse

	if err := c.baseClient.Get(GetCategoriesListRequestPoint, nil, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *ClientImpl) GetCategoriesData(request *GetCategoriesDataRequest) (*GetCategoriesDataResponse, error) {
	if err := request.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	var response GetCategoriesDataResponse

	opts := &base.RequestOptions{
		QueryParams: map[string]string{
			"order":                   request.Order,
			"per_page":                fmt.Sprintf("%d", request.PerPage),
			"page":                    fmt.Sprintf("%d", request.Page),
			"sparkline":               fmt.Sprintf("%v", request.Sparkline),
			"price_change_percentage": request.PriceChangePercentage,
		},
	}

	if err := c.baseClient.Get(GetCategoriesDataRequestPoint, opts, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
