# globalPath

This package allows a global filepath to be initialized, which can then be used from other packages.

## Support

If you use or find value in this package, please consider donating at PayPal: [https://www.paypal.me/jismithpp](https://www.paypal.me/jismithpp)

## Example

```golang
package main

import (
	"fmt"

	"github.com/theTardigrade/globalPath"
)

func main() {
	// initializes the base path as the path to the "models" subdirectory of the
	// "data" subdirectory of the directory where this package is located
	globalPath.Init("data", "models")

	// prints the base path
	fmt.Println(globalPath.Get())

	// prints the path to the "z" subdirectory of the "y" subdirectory of the
	// "x" subdirectory of the base path
	fmt.Println(globalPath.Join("x", "y", ".", "z"))

	// reinitializes the base path to the directory where this package is located
	globalPath.Init()

	// prints path to the directory above the base path
	fmt.Println(globalPath.Join(".."))
}
```