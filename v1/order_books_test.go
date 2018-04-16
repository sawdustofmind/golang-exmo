package exmo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestOrderBooksServiceGet(t *testing.T) {
	GetTradesHandler := func(w http.ResponseWriter, r *http.Request) {
		resp := `{
							"BTC_USD": {
								"ask_quantity": "3",
								"ask_amount": "500",
								"ask_top": "100",
								"bid_quantity": "1",
								"bid_amount": "99",
								"bid_top": "99",
								"ask": [[100,1,100],[200,2,400]],
								"bid": [[99,1,99]]
							}
						}`
		w.Write([]byte(resp))
	}

	server := httptest.NewServer(http.HandlerFunc(GetTradesHandler))
	defer server.Close()

	client := NewClient()
	client.BaseURL, _ = url.Parse(server.URL)

	orderBooks, err := client.OrderBooks.Get([]string{"BTC_USD"})

	if err != nil {
		t.Error(err)
	}
	if orderBooks == nil {
		t.Error("Expected", 1)
	}
}
