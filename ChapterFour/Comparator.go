package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var number int
	fmt.Scan(&number)

	fmt.Println("Enter another number: ")
	var numberTwo int
	fmt.Scan(&numberTwo)

	if number > numberTwo {
		fmt.Println("1")
	} else if number < numberTwo {
		fmt.Println("-1")
	} else {
		fmt.Println("0")
	}
}
