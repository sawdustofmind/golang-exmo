package exmo

import (
	"net/url"
	"strings"
)

type TradesService struct {
	c *Client
}

type TradeStruct struct {
	TradeID  int64 `json:"trade_id,int"`
	Type     string
	Price    string
	Quantity string
	Amount   string
	Date     int64
}

type Trades map[string][]TradeStruct

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
