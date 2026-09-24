package main

import "fmt"

func main() {
	pi := 0.0
	divisor := 1.0

	for count := 1; count <= 200000; count++ {
		if count%2 != 0 {
			pi += 4.0 / divisor
		} else {
			pi -= 4.0 / divisor
		}
		divisor += 2
	}
	fmt.Println(pi)

}
