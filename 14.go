// Longest Collatz sequence
package main

import "fmt"

func number_of_terms(i int64) int32 {
	var counter int32 = 1
	for i > 1 {

		if i%2 == 0 {
			i = i / 2
			counter++
		} else if i%2 != 0 {
			i = 3*i + 1
			counter++
		}
		if i == 1 {
			break
		}
	}
	return counter
}
func main() {
	var max int32 = 0
	for j := 999999; j > 1; j-- {

		if number_of_terms(int64(j)) > max {
			max = number_of_terms(int64(j))
			fmt.Println(j)
		}

	}
}
