package exmo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestTradesServiceGet(t *testing.T) {
	GetOrderBooksHandler := func(w http.ResponseWriter, r *http.Request) {
		resp := `{
			"BTC_USD": [
				{
					"trade_id": 3,
					"type": "sell",
					"price": "100",
					"quantity": "1",
					"amount": "100",
					"date": 1435488248
				}
			]
		}`
		w.Write([]byte(resp))
	}

	server := httptest.NewServer(http.HandlerFunc(GetOrderBooksHandler))
	defer server.Close()

	client := NewClient()
	client.BaseURL, _ = url.Parse(server.URL)

	trades, err := client.Trades.Get([]string{"BTC_USD"})

	if err != nil {
		t.Error(err)
	}
	if trades["BTC_USD"] == nil {
		t.Error("Expected", 1)
	}
}
