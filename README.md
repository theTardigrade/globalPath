# golang-globalFilepath

This package allows a global filepath to be initialized, which can then be used from other packages.

## Example

```golang
package main

import (
	"fmt"

	gf "github.com/theTardigrade/golang-globalFilepath"
)

func main() {
	// initializes the base path as the path to the "models" subdirectory of the
	// "data" subdirectory of the directory where this package is located
	gf.Init("data", "models")

	// prints the base path
	fmt.Println(gf.Get())

	// prints the path to the "z" subdirectory of the "y" subdirectory of the
	// "x" subdirectory of the base path
	fmt.Println(gf.Join("x", "y", ".", "z"))

	// reinitializes the base path to the directory where this package is located
	gf.Init()

	// prints path to the directory above the base path
	fmt.Println(gf.Join(".."))
}
```