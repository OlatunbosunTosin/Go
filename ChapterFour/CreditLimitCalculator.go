package main

import "fmt"

func main() {
	fmt.Print("Enter balance at beginning of the month: ")
	var balance int
	fmt.Scan(&balance)

	fmt.Print("Enter total charges for the month: ")
	var totalCharges int
	fmt.Scan(&totalCharges)

	fmt.Print("Enter total credits for the month: ")
	var totalCredits int
	fmt.Scan(&totalCredits)

	fmt.Print("Enter allowed credit limit: ")
	var creditLimit int
	fmt.Scan(&creditLimit)

	newBalance := balance + totalCharges - totalCredits
	if newBalance > creditLimit {
		fmt.Println("Credit limit exceeded")
	} else {
		fmt.Print("Credit limit not exceeded")
	}

}
