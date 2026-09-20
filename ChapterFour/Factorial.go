package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var number int
	fmt.Scan(&number)
	factorial := 1

	if number < 0 {
		fmt.Println("Number is negative")
	} else if number == 0 {
		fmt.Println("1")
	} else {
		for index := number; index >= 1; index-- {
			factorial *= index
		}
		fmt.Println(factorial)
	}
}
