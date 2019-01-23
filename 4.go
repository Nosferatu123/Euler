// Largest palindrome product
package main

import (
	"fmt"
)

func main() {
	var a, b, c, d, e, f int32 = 0, 0, 0, 0, 0, 0
	var arr []int32
	var counter = 1
	arr = append(arr, 999999)
	for i := 999998; i > 100000; i-- {
		a = int32(int32(i) / 100000)
		b = int32((int32(i) - a*100000) / 10000)
		c = int32((int32(i) - a*100000 - b*10000) / 1000)
		d = int32((int32(i) - a*100000 - b*10000 - c*1000) / 100)
		e = int32((int32(i) - a*100000 - b*10000 - c*1000 - d*100) / 10)
		f = int32(int32(i) - a*100000 - b*10000 - c*1000 - d*100 - e*10)
		if a == f && b == e && c == d {
			arr = append(arr, int32(i))
		}
	}
	for j := 1; j < len(arr); j++ {
		if counter == 0 {
			break
		}
		for k := 100; k <= 999; k++ {
			if arr[j]%int32(k) == 0 && arr[j]/int32(k) < 1000 {
				fmt.Println(arr[j], ",", int32(k), "- first 3-digit number", ",", arr[j]/int32(k), "- second 3-digit number")
				counter = 0
				break
			}

		}
	}
}
