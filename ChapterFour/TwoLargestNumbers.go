package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var number int
	fmt.Scan(&number)
	largestNumber := 0
	secondLargestNumber := 0

	for counter := 1; counter < 10; counter++ {
		if number > largestNumber {
			secondLargestNumber = largestNumber
			largestNumber = number
		} else if number > secondLargestNumber {
			secondLargestNumber = number
		}
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
		for number == largestNumber || number == secondLargestNumber {
			fmt.Print("Enter another number: ")
			fmt.Scan(&number)
		}
	}
	fmt.Printf("Largest number = %d\nSecond largest number = %d\n", largestNumber, secondLargestNumber)
}
