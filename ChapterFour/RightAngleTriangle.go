package main

import "fmt"

func main() {
	fmt.Print("Enter triangle base number: ")
	var triangleNumber int
	fmt.Scan(&triangleNumber)

	for row := 1; row <= triangleNumber; row++ {
		for column := 1; column <= row; column++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
