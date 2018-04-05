package exmo

import (
	"errors"
	"net/url"
	"strconv"
)

type OrderService struct {
	c *Client
}

type OrderCreateResponse struct {
	Result  bool
	Error   string
	OrderID int64 `json:"order_id,int"`
}

type OrderCancelResponse struct {
	Result bool
	Error  string
}

type OrderTradesResponse struct {
	Type        string
	InCurrency  string `json:"in_currency"`
	OutCurrency string `json:"out_currency"`
	OutAmount   string `json:"out_amount"`
	Trades      []struct {
		TradeID  int64 `json:"trade_id,int"`
		Type     string
		OrderID  int64 `json:"order_id,int"`
		Pair     string
		Price    string
		Quantity string
		Amount   string
		Date     int64
	}
}

//Create OrderService creates order and return orderId
func (a *OrderService) Create(pair string, quantity float64, price float64, orderType string) (OrderCreateResponse, error) {
	params := url.Values{}

	params.Add("pair", pair)
	params.Add("quantity", strconv.FormatFloat(quantity, 'f', 6, 64))
	params.Add("price", strconv.FormatFloat(price, 'f', 6, 64))
	params.Add("type", orderType)

	req, err := a.c.newAuthenticatedRequest("order_create", params)

	if err != nil {
		return OrderCreateResponse{}, err
	}

	var v = OrderCreateResponse{}

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return OrderCreateResponse{}, err
	}

	return v, nil
}

//Trades OrderService gives  order trades
func (a *OrderService) Trades(orderID int64) (OrderTradesResponse, error) {
	params := url.Values{}

	params.Add("order_id", strconv.Itoa(int(orderID)))

	req, err := a.c.newAuthenticatedRequest("order_trades", params)

	if err != nil {
		return OrderTradesResponse{}, err
	}

	var v OrderTradesResponse

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return OrderTradesResponse{}, err
	}

	return v, nil
}

//Cancel OrderService cancel order by given id
func (a *OrderService) Cancel(orderID int64) error {
	params := url.Values{}

	params.Add("order_id", strconv.Itoa(int(orderID)))

	req, err := a.c.newAuthenticatedRequest("order_cancel", params)

	if err != nil {
		return err
	}

	var v OrderCancelResponse

	_, err = a.c.performRequest(req, &v)

	if err != nil {
		return err
	}

	if !v.Result {
		return errors.New(v.Error)
	}

	return nil
}
