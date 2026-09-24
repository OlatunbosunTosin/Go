package main

import (
	"fmt"
	"strings"
)

func main() {

	countA := 0
	countB := 0
	countC := 0
	countD := 0
	countE := 0
	for count := 0; count < 5; count++ {
		fmt.Print("Enter student name: ")
		var name string
		fmt.Scan(&name)
		fmt.Print("Enter student grade: ")
		var grade string
		fmt.Scan(&grade)

		grade = strings.ToLower(grade)
		switch grade {
		case "a":
			countA++
		case "b":
			countB++
		case "c":
			countC++
		case "d":
			countD++
		case "e":
			countE++
		default:
			fmt.Print("invalid")
		}
	}
	fmt.Printf("%d%d%d%d%d", countA, countB, countC, countD, countE)
}
