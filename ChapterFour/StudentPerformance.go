package main

import "fmt"

func main() {
	passes := 0
	failures := 0
	studentCount := 0

	for studentCount < 10 {
		fmt.Println("Enter result (1 = pass, 2 = fail): ")
		var result int
		fmt.Scan(&result)

		for result != 1 && result != 2 {
			fmt.Println("invalid (1 = pass, 2 = fail): ")
			fmt.Scan(&result)
		}
		if result == 1 {
			passes = passes + 1
		} else {
			failures = failures + 1
		}

		studentCount++
	}
	fmt.Printf("Passed: %d\nFailed: %d\n", passes, failures)
	if passes > 8 {
		fmt.Print("Bonus to instructor!")
	}
}
