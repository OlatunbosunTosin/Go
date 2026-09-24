package main

import "fmt"

func main() {
	sum := 0
	for count := 1; count <= 30; count++ {
		if count%3 == 0 {
			sum += count
		}
	}
	fmt.Println(sum)
}
