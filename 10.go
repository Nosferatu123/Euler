// Summation of primes

package main

import (
	"fmt"
	"math"
)

func prime_or_not(n int) bool {

	if n == 2 {
		return true
	}

	if n%2 == 0 && n > 2 {
		return false
	}

	var sqrt = int(math.Sqrt(float64(n)))

	for i := 3; i <= sqrt+1; i += 2 {

		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	var sum int = 0

	for j := 2; int64(j) < 2000000; j++ {

		if prime_or_not(j) {
			sum = sum + j
		}

	}
	fmt.Println(sum)
}
