package exmo

type OrderService struct {
	c *Client
}

//Create OrderService creates order and return orderId
func (a *OrderService) Create(pair string, quantity float64, price float64, orderType string) (int64, error) {
	return 1, nil
}
