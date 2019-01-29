//Power digit sum
package main

import (
	"fmt"
	"math/big"
)

func main() {
	sum := big.NewInt(0)
	ten := big.NewInt(10)
	counter := 0
	n := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil)
	for counter <= 1000 {

		sum.Add(sum, new(big.Int).Mod(n, ten))
		n.Div(n, big.NewInt(10)) // n/10
		counter++

	}
	fmt.Println(sum)
}
