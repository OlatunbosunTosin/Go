package main

import "fmt"

func main() {
	x := 9
	y := 11
	//if x > 5 {
	//	if y > 5 {
	//		fmt.Println("x and y are > 5")
	//	}
	//} else {
	//	fmt.Println("x is <= 5")
	//}

	if x < 10 {
		if y > 10 {
			fmt.Println("*****")
		}
	} else {
		fmt.Println("#####")
		fmt.Println("$$$$$")
	}
}
