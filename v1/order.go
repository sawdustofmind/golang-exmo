package exmo

import (
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
