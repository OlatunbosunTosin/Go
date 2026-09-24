package main

import "fmt"

func main() {
	for row := 1; row <= 10; row++ {
		for column := 1; column <= row; column++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()
	for row := 1; row <= 10; row++ {
		for column := row; column <= 10; column++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()
	for row := 1; row <= 10; row++ {
		for column := 2; column <= row; column++ {
			fmt.Print(" ")
		}
		for space := 10; space >= row; space-- {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Println()
	for row := 1; row <= 10; row++ {
		for column := 9; column >= row; column-- {
			fmt.Print(" ")
		}
		for space := 1; space <= row; space++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
