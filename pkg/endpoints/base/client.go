package base

import (
	"fmt"
	"time"

	"resty.dev/v3"
)

// ClientOption defines client configuration options
type ClientOption func(*resty.Client)

// Config defines the base client configuration
type Config struct {
	// BaseURL is the base URL for the API
	BaseURL string
	// Timeout is the request timeout duration
	Timeout time.Duration
	// RetryCount is the number of retry attempts
	RetryCount int
	// RetryWaitTime is the duration to wait between retries
	RetryWaitTime time.Duration
	// APIKey is the API key for authentication
	APIKey string
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		BaseURL:       "https://api.coingecko.com/api/v3",
		Timeout:       10 * time.Second,
		RetryCount:    3,
		RetryWaitTime: 1 * time.Second,
	}
}

// BaseClient is the base client structure
type BaseClient struct {
	httpClient *resty.Client
	config     *Config
}

// NewBaseClient creates a new base client
func NewBaseClient(config *Config, options ...ClientOption) *BaseClient {
	if config == nil {
		config = DefaultConfig()
	}

	client := resty.New()
	client.SetBaseURL(config.BaseURL)
	client.SetTimeout(config.Timeout)

	// Configure retries
	client.SetRetryCount(config.RetryCount)
	client.SetRetryWaitTime(config.RetryWaitTime)

	// Set API key
	if config.APIKey != "" {
		client.SetHeader("x-cg-pro-api-key", config.APIKey)
	}

	// Apply custom options
	for _, option := range options {
		option(client)
	}

	return &BaseClient{
		httpClient: client,
		config:     config,
	}
}

// RequestOptions defines request options
type RequestOptions struct {
	// PathParams are path parameters
	PathParams map[string]string
	// QueryParams are query parameters
	QueryParams map[string]string
	// Body is the request body
	Body interface{}
	// Headers are request headers
	Headers map[string]string
}

// DoRequest executes an HTTP request
func (c *BaseClient) DoRequest(method, path string, opts *RequestOptions, result interface{}) error {
	if opts == nil {
		opts = &RequestOptions{}
	}

	req := c.httpClient.R()

	// Set path parameters
	if len(opts.PathParams) > 0 {
		req.SetPathParams(opts.PathParams)
	}

	// Set query parameters
	if len(opts.QueryParams) > 0 {
		req.SetQueryParams(opts.QueryParams)
	}

	// Set request body
	if opts.Body != nil {
		req.SetBody(opts.Body)
	}

	// Set request headers
	if len(opts.Headers) > 0 {
		req.SetHeaders(opts.Headers)
	}

	// Set response
	if result != nil {
		req.SetResult(result)
	}

	// Execute request
	res, err := req.Execute(method, path)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}

	// Check status code
	if res.StatusCode() >= 300 {
		return fmt.Errorf("API error: status code %d, response: %s", res.StatusCode(), res.String())
	}

	return nil
}

// Get executes a GET request
func (c *BaseClient) Get(path string, opts *RequestOptions, result interface{}) error {
	return c.DoRequest("GET", path, opts, result)
}

// Post executes a POST request
func (c *BaseClient) Post(path string, opts *RequestOptions, result interface{}) error {
	return c.DoRequest("POST", path, opts, result)
}

// Put executes a PUT request
func (c *BaseClient) Put(path string, opts *RequestOptions, result interface{}) error {
	return c.DoRequest("PUT", path, opts, result)
}

// Delete executes a DELETE request
func (c *BaseClient) Delete(path string, opts *RequestOptions, result interface{}) error {
	return c.DoRequest("DELETE", path, opts, result)
}
