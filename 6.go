// Sum square difference
package main

import (
	"fmt"
)

func main() {
	var sum_of_squares, sum int32 = 0, 0
	for i := 1; i <= 100; i++ {
		sum_of_squares = sum_of_squares + int32(i*i)
		sum = sum + int32(i)
	}
	fmt.Println(sum*sum - sum_of_squares)

}
