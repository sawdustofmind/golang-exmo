package exmo

import (
	"net/url"
)

type TickerService struct {
	c *Client
}

type TickerStruct struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	High      string
	Low       string
	Avg       string
	Vol       string
	VolCurr   string `json:"vol_curr"`
	Updated   int64
}

type Ticker map[string]TickerStruct

//Get TickerService returns ticker
func (a *TickerService) Get() (Ticker, error) {
	req, err := a.c.newRequest("GET", "ticker", url.Values{})

	if err != nil {
		return Ticker{}, err
	}

	var v Ticker

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return Ticker{}, err
	}

	return v, nil
}
