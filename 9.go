//Special Pythagorean triplet
package main

import (
	"fmt"
)

func main() {
	var c int32
	var counter int32 = 0
	for a := 1; a < 1000; a++ {
		if counter == 1 {
			break
		}
		for b := 1; b < 1000; b++ {
			c = int32(1000 - a - b)
			if int32(a*a+b*b) == c*c {
				fmt.Println(int32(a) * int32(b) * c)
				counter = 1
			}
		}

	}
}
