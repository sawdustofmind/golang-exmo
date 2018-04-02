package tests

import "testing"

func TestUserInfo(t *testing.T) {
	_, err := client.User.Info()

	if err.Error() != "POST https://api.bitfinex.com/v1/order/new: 400 Invalid order: not enough exchange balance for 1.0 BTCUSD at 299.0" {
		t.Fatalf("OrderBook.Get() returned error: %v", err)
	}
}
