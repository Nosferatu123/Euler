// 10001st prime
package main

import (
	"fmt"
)

func () {

	var arr []int32
	var counter int32
	var sum int32
	arr = append(arr, 2)
	for i := 3; i <= 10000; i++ {
		counter = 0
		for j := 2; int32(j) < int32(i/2)+1; j++ {
			if i%j == 0 {
				counter++
			}
		}
		if counter == 0 {
			arr = append(arr, int32(i))
			sum++
		}
	}
	fmt.Println(sum)
}
