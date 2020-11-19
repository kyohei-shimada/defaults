# defaults

This package provides shorthands to initialize with default value.

## Usage

`defaults.{Type}(arg1, arg2, ...)`  return the first non-zero argument.

```go
package main

import (
	"fmt"

	"github.com/kyohei-shimada/defaults"
)

func main() {
	// Get a first non empty string
	fmt.Println("%s\n", defaults.String("foo", "bar")) // "foo"
	fmt.Println("%s\n", defaults.String("", "foo", "bar"))      // "bar"

	// Get a first non zero value
	fmt.Printf("%d\n", defaults.Int32(1, 10, 100, 1000)) // 1
	fmt.Printf("%d\n", defaults.Int32(0, 123, 0, 999)) // 123
}
```

## Typically Usage

Typically, `defaults.{Type}(arg1, arg2, ...)` use initialization.

```go
package main

import (
	"os"

	"github.com/kyohei-shimada/defaults"
)

type ExternalApiClient struct {
	Host string
	AccessKey string
}

func main() {
	conf := // Load some config from file

	client := &ExternalApiClient{
		// configuration setttings take precedence in the following order.
		// 1. Environment variables
		// 2. loaded config from file
		// (3. default value)
		Host: defaults.String(os.Getenv("HOST"), conf.Host, "localhost"),
		AccessKey: defaults.Int(os.Getenv("ACCESS_KEY"), conf.AccessKey),
	}
}
```