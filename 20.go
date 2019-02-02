//Factorial digit sum
package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(1)
	j := big.NewInt(0)
	for i := 1; i <= 100; i++ {
		j.Add(j, big.NewInt(1))
		n.Mul(n, j)
	}
	fmt.Println(n)
	sum := big.NewInt(0)
	ten := big.NewInt(10)
	zero := big.NewInt(0)

	for n.Cmp(zero) > 0 {
		sum.Add(sum, new(big.Int).Mod(n, ten))
		n.Div(n, big.NewInt(10)) // n/10

	}
	fmt.Println(sum)
}
