package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("How many numbers are you collecting? ")
	var number int
	fmt.Scan(&number)

	minimum := math.MaxInt32
	maximum := -1
	for index := 1; index <= number; index++ {
		fmt.Print("Enter a number: ")
		var input int
		fmt.Scan(&input)
		if input < minimum {
			minimum = input
		}
		if input > maximum {
			maximum = input
		}
	}
	sum := minimum + maximum
	fmt.Println("The sum is", sum)
}
