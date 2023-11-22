# go-helpers
A Golang library for common helper function that I re-implement in many projects, but now in a single place.

## How to install the package
Run: `$ go get github.com/officer47p/go-helpers`

## How to use the package
```go
package main

import (
	"fmt"

	gohelpers "github.com/officer47p/go-helpers"
)

func main() {
	s := gohelpers.FloatToString(0.022)

	fmt.Println(s)
}
```
