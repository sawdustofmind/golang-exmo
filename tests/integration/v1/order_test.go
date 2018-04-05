package tests

// import (
// 	"testing"
// )

// func TestOrderCreateAndCancel(t *testing.T) {
// 	order, err := client.Order.Create("ETH_UAH", 0.02, 1000.0, "buy")

// 	if err != nil {
// 		t.Fatalf("Order.Create() returned error: %v", err)
// 	}

// 	t.Logf("Order created: %d", order.OrderID)

// 	_, err = client.Order.Trades(order.OrderID)

// 	if err != nil {
// 		t.Logf("Order id: %d is not canceled!", order.OrderID)
// 		t.Fatalf("Order.Trades() returned error: %v", err)
// 	}

// 	err = client.Order.Cancel(order.OrderID)

// 	if err != nil {
// 		t.Fatalf("Order.Cancel() returned error: %v", err)
// 	}
// }
