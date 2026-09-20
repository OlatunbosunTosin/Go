package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var number int
	fmt.Scan(&number)
	initialNumber := number
	reversedNumber := 0

	for number < 10000 {
		fmt.Println("Enter a number: ")
		fmt.Scan(&number)
	}
	for number > 0 {
		remainder := number % 10
		reversedNumber = reversedNumber*10 + remainder
		number = number / 10
	}
	if initialNumber == reversedNumber {
		fmt.Println("It is a palindrome number")
	} else {
		fmt.Println("It is not a palindrome number")
	}

}
