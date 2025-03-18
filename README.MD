# Golang options pattern implementation

## Installation

```shell
go get github.com/go-gang/options
```

## Usage
```go
package main

import (
    "github.com/go-gang/options"
)

type config struct {
	host string
	port int
}

func WithHost(host string) options.Option[config] {
	return options.Func(func(c *config) {
        c.host = host
	})
}

func WithPort(port int) options.Option[config] {
	return options.Func(func(c *config) {
		c.port = port
	})
}

func main() {
	fromZeroValues := options.New(WithHost("localhost"), WithPort(80))
	
	...
	
	fromDefaultValues := config{
		host: "localhost",
		port: 80,
    }
	
	options.Apply(&fromDefaultValues, WithPort(8080))
	
	...
}

```

## License

The project is licensed under the [MIT License](LICENSE).