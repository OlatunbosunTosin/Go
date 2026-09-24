package main

import "fmt"

func main() {
	fmt.Print("Enter n: ")
	var n int
	fmt.Scan(&n)
	for n < 1 || n > 100 {
		fmt.Print("Enter n: ")
		fmt.Scan(&n)
	}

	sum := 0
	for count := 1; count <= n; count++ {
		sum += count
	}
	fmt.Printf("%d\t%d", n, sum)
}
