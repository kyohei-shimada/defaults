# defaults

This package provides shorthands to initialize with default value.

## Usage

```go
package main

import (
	"fmt"

	"github.com/kyohei-shimada/defaults"
)

func main() {

	fmt.Printf("%s\n", defaults.String("Hello", "default value")) // "Hello"
	fmt.Printf("%s\n", defaults.String("", "default value"))      // "default value"

	fmt.Printf("%d\n", defaults.Int32(1, 999)) // 1
	fmt.Printf("%d\n", defaults.Int32(0, 999)) // 999
}
```
