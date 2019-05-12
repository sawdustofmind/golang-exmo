package exmo

import (
	"net/url"
)

// Client manages all the communication with the API.
type Client struct {
	// Base URL for API requests.
	BaseURL *url.URL

	APIKey    string
	APISecret string

	Trades      *TradesService
	OrderBooks  *OrderBooksService
	Order       *OrderService
	User        *UserService
	Ticker      *TickerService
	PairSetting *PairSettingService
}

// NewClient creates new API client.
func NewClient() *Client {
	baseURL, _ := url.Parse(BaseURL)

	c := &Client{BaseURL: baseURL}

	c.Trades = &TradesService{c: c}
	c.OrderBooks = &OrderBooksService{c: c}
	c.Order = &OrderService{c: c}
	c.User = &UserService{c: c}
	c.Ticker = &TickerService{c: c}
	c.PairSetting = &PairSettingService{c: c}

	return c
}

// Auth sets api key and secret for usage is requests that requires authentication.
func (c *Client) Auth(key string, secret string) *Client {
	c.APIKey = key
	c.APISecret = secret

	return c
}
