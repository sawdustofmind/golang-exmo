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

```

### Authentication

``` go

```

### Order create

``` go

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
