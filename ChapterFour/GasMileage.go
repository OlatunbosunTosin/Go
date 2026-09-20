package main

import "fmt"

func main() {
	fmt.Print("Enter miles driven: ")
	var miles float64
	fmt.Scan(&miles)
	count := 0
	sum := 0.0
	averageMilesPerGallon := 0.0

	for miles != -1 {
		fmt.Print("Enter gallons used: ")
		var gallonsUsed int
		fmt.Scan(&gallonsUsed)
		count++

		milesPerGallon := (miles) / float64(gallonsUsed)
		sum += milesPerGallon
		averageMilesPerGallon = sum / float64(count)
		fmt.Print("Enter miles driven: ")
		fmt.Scan(&miles)
	}
	fmt.Printf("Total miles per gallon = %f\n", sum)
	fmt.Printf("Average miles = %f", averageMilesPerGallon)
}
