//Smallest multiple
package main

import "fmt"

func main() {
	var n int64 = 20
	sum := n * (n - 1)
	for i := 20; i > 10; i-- {
		if sum%int64(i) != 0 {
			sum = sum + n*(n-1)
			i = 20
		}
	}
	fmt.Println(sum)
}
