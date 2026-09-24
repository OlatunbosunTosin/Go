package main

import (
	"fmt"
	"strings"
)

func main() {
	chart := ``
	for count := 0; count < 5; count++ {
		fmt.Printf("Enter a number: ")
		var number int
		fmt.Scan(&number)
		for number < 1 || number > 30 {
			fmt.Printf("Enter number between 1 and 30: ")
			fmt.Scan(&number)
		}
		chart += strings.Repeat("*", number)
		chart += "\n"
	}
	fmt.Println(chart)
}
