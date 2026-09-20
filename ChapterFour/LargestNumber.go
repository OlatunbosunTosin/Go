package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scan(&number)
	largestNumber := 0

	for counter := 1; counter < 10; counter++ {
		if number > largestNumber {
			largestNumber = number
		}
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
	}
	fmt.Printf("Largest number = %d\n", largestNumber)
}
