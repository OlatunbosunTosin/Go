package main

import "fmt"

func main() {
	x := 7
	y := 3
	x = y
	y += 1
	y += 1
	x = y
	fmt.Print(x)
}
