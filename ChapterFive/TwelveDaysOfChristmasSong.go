package main

import "fmt"

func main() {

	for count := 1; count <= 12; count++ {
		fmt.Printf("On the ")
		switch count {
		case 1:
			fmt.Print("first")
		case 2:
			fmt.Print("second")
		case 3:
			fmt.Print("third")
		case 4:
			fmt.Print("fourth")
		case 5:
			fmt.Print("fifth")
		case 6:
			fmt.Print("sixth")
		case 7:
			fmt.Print("seventh")
		case 8:
			fmt.Print("eighth")
		case 9:
			fmt.Print("ninth")
		case 10:
			fmt.Print("tenth")
		case 11:
			fmt.Print("eleventh")
		case 12:
			fmt.Print("twelfth")
		}
		fmt.Printf(" day of Christmas my true love gave to me ")

		for day := count; day >= 1; day-- {
			switch day {
			case 1:

				if count == 1 {
					fmt.Print("A partridge in a pear tree")
				} else {
					fmt.Print("and a partridge in a pear tree")
				}

			case 2:
				fmt.Print("Two turtle doves")
			case 3:
				fmt.Print("Three French hens")
			case 4:
				fmt.Print("Four calling birds")
			case 5:
				fmt.Print("Five gold rings")
			case 6:
				fmt.Print("Six geese a-laying")
			case 7:
				fmt.Print("Seven swans a-swimming")
			case 8:
				fmt.Print("Eight maids a-milking")
			case 9:
				fmt.Print("Nine ladies dancing")
			case 10:
				fmt.Print("Ten lords a-leaping")
			case 11:
				fmt.Print("Eleven pipers piping")
			case 12:
				fmt.Print("Twelve drummers drumming")
			}
			fmt.Println()
		}
	}
}
