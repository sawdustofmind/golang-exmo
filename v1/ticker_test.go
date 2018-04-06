package exmo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestTickerServiceGet(t *testing.T) {
	GetTickerHandler := func(w http.ResponseWriter, r *http.Request) {
		resp := `{
							"BTC_USD": {
								"buy_price": "589.06",
								"sell_price": "592",
								"last_trade": "591.221",
								"high": "602.082",
								"low": "584.51011695",
								"avg": "591.14698808",
								"vol": "167.59763535",
								"vol_curr": "99095.17162071",
								"updated": 1470250973
							}
						}`

		w.Write([]byte(resp))
	}

	server := httptest.NewServer(http.HandlerFunc(GetTickerHandler))
	defer server.Close()

	client := NewClient()
	client.BaseURL, _ = url.Parse(server.URL)

	_, err := client.Ticker.Get()

	if err != nil {
		t.Error(err)
	}
}
