package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	amount := 0.0
	fmt.Printf("%s%20s\n", "Year", "Amount on deposit")
	for count := 5; count <= 10; count++ {
		rate := float64(count) / 100.
		fmt.Printf("%.2f Interest\n", rate)
		for year := 1; year <= 10; year++ {
			amount = principal * math.Pow((1.0+rate), float64(year))
			fmt.Printf("%4d%20.2f\n", year, amount)
		}
		fmt.Println()
	}
}
