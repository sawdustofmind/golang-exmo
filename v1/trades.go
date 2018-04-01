package exmo

import (
	"net/url"
	"strings"
	"time"
)

type TradesService struct {
	c *Client
}

type TradeStruct struct {
	TradeID  int       `json:"trade_id"`
	Type     string    `json:"type"`
	Price    float64   `json:"price"`
	Quantity float64   `json:"quantity"`
	Amount   float64   `json:"amount"`
	Date     time.Time `json:"date"`
}

type Trades map[string][]*TradeStruct

//Get TradesService returns trades for selected pairs
func (a *TradesService) Get(pairs []string) (Trades, error) {
	params := url.Values{}

	params.Add("pair", strings.Join(pairs, ","))
	req, err := a.c.newRequest("GET", "trades", params)

	if err != nil {
		return Trades{}, err
	}

	var v Trades
	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return Trades{}, err
	}

	return v, nil
}
