// Largest prime factor
package main

import (
	"fmt"
)

func Largest_prime_factor() {
	var n int64 = 600851475143
	var counter int32 = 0
	var arr []int64
	arr = append(arr, 2)
	for i := int64(math.Sqrt(float64(n)) + 1); i > 2; i-- {
		counter = 0
		if n%i == 0 {
			for j := 2; int64(j) < int64(i/2+1); j++ {
				if i%int64(j) == 0 {
					counter++
				}
			}
		}
		if counter == 0 {
			arr = append(arr, int64(i))
		}
	}
	for k := 1; k < len(arr); k++ {
		if n%int64(arr[k]) == 0 {
			fmt.Println(arr[k])
			break
		}
	}
}
