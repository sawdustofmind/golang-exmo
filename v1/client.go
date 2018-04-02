package exmo

import (
	"encoding/json"
	"net/url"
)

// Client manages all the communication with the API.
type Client struct {
	// Base URL for API requests.
	BaseURL *url.URL

	APIKey    string
	APISecret string

	Trades     *TradesService
	OrderBooks *OrderBooksService
	Order      *OrderService
	User       *UserService
}

// NewClient creates new API client.
func NewClient() *Client {
	baseURL, _ := url.Parse(BaseURL)

	c := &Client{BaseURL: baseURL}

	c.Trades = &TradesService{c: c}
	c.OrderBooks = &OrderBooksService{c: c}
	c.Order = &OrderService{c: c}
	c.User = &UserService{c: c}

	return c
}

// Auth sets api key and secret for usage is requests that requires authentication.
func (c *Client) Auth(key string, secret string) *Client {
	c.APIKey = key
	c.APISecret = secret

	return c
}

// checkResponse checks response status code and response
// for errors.
func checkResponse(r *Response) error {
	if c := r.Response.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	// Try to decode error message
	errorResponse := &ErrorResponse{Response: r}
	err := json.Unmarshal(r.Body, errorResponse)
	if err != nil {
		errorResponse.Message = "Error decoding response error message. " +
			"Please see response body for more information."
	}

	return errorResponse
}
