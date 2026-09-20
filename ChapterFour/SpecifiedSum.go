package main

import "fmt"

func main() {
	fmt.Println("Enter one number: ")
	var number int
	fmt.Scan(&number)
	total := 0

	for total < number {
		fmt.Println("Enter a number: ")
		var integer int
		fmt.Scan(&integer)
		total += integer
	}
	fmt.Println(total)
}
