package main

import (
	"fmt"
	"math"
)

func main() {
	for count := 1; count <= 5; count++ {
		for counter := 1; counter <= 4; counter++ {
			fmt.Print(math.Pow(float64(count), float64(counter)))
			fmt.Print("\t")
		}
		fmt.Print("\n")

	}
}
