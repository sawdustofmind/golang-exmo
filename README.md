# Exmo Trading API for Golang.

## Installation

``` bash
go get github.com/asxcandrew/golang-exmo
```

## Usage

### Basic requests

``` go
package main

import (
	"fmt"
	"github.com/asxcandrew/golang-exmo/v1"
)

client := exmo.NewClient()
	pairs := []string{"BTC_USD", "ETC_USD"}

	trades, err := client.Trades.Get(pairs)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(trades)
	}
```

### Authentication

``` go
client := exmo.NewClient().Auth(key, secret)

```

### Order create

``` go
client := exmo.NewClient().Auth(key, secret)

order, err := client.Order.Create("ETH_UAH", 0.02, 1000.0, "buy")

if err == nil {
	fmt.Println(order)
}
```

See [examples](https://github.com/asxcandrew/golang-exmo/tree/master/examples)

## Testing

All integration tests are stored in `tests/integration` directory.

Run tests using:
``` bash
export EXMO_API_KEY="api-key"
export EXMO_API_SECRET="api-secret"
go test -v ./tests/integration
```

## Contributing

1. Fork it (https://github.com/asxcandrew/golang-exmo/fork)
2. Create your feature branch (`git checkout -b my-new-feature)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
