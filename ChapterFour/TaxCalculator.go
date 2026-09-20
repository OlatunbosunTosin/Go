package main

import "fmt"

func main() {
	fmt.Print("What is your earning: ")
	var earning float64
	fmt.Scan(&earning)
	tax := 0.0

	if earning == 30000 {
		tax = (15.0 / 100) * earning
		fmt.Printf("Tax = %.2f", tax)
	} else if earning > 30000 {
		tax = (20.0 / 100) * earning
		fmt.Printf("Tax = %.2f", tax)
	}
}
