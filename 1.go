//Multiples of 3 and 5
package main

func main() {
	var sum_1_task int = 0
	for i := 1; int32(i) < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum_1_task = sum_1_task + i
		}
	}
}
