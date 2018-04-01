package exmo

import (
	"net/url"
	"strings"
)

type OrderBooksService struct {
	c *Client
}

type OrderBookStruct struct {
	AskQuantity string `json:"ask_quantity"`
	AskAmount   string `json:"ask_amount"`
	AskTop      string `json:"ask_top"`
	BidQuantity string `json:"bid_quantity"`
	BidAmount   string `json:"bid_amount"`
	BidTop      string `json:"bid_top"`
	Ask         [][]int
	Bid         [][]int
}

type OrderBooks map[string]OrderBookStruct

//Get OrderBooksService returns order books for selected pairs
func (a *OrderBooksService) Get(pairs []string) (OrderBooks, error) {
	params := url.Values{}
	params.Add("pair", strings.Join(pairs, ","))

	req, err := a.c.newRequest("GET", "order_book", params)

	if err != nil {
		return OrderBooks{}, err
	}

	var v OrderBooks
	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return OrderBooks{}, err
	}

	return v, nil
}
