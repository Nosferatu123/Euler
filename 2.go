// Even Fibonacci numbers
package main

import (
	"fmt"
)

func Even_Fibonacci_numbers() {
	var arr []int64
	var sum_2_task int64
	arr = append(arr, 1)
	arr = append(arr, 2)
	for i := 0; int64(i) < 10000; i++ {

		if arr[i] >= 4000000 {
			break
		}
		arr = append(arr, arr[int64(i)]+arr[int64(i+1)])

	}
	for j := 0; j < len(arr); j++ {
		if arr[j]%2 == 0 {
			sum_2_task = sum_2_task + arr[j]
		}
	}
	fmt.Println(arr)
	fmt.Println(sum_2_task)

}
