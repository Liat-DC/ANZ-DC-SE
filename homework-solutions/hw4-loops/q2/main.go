// This program reads five numbers from the user and prints the largest number entered using loops

package main

import (
	"fmt"
)

func main() {

	var num, max int

	fmt.Println("Please enter 10 numbers")
	fmt.Print(" Num 1: ")
	fmt.Scan(&num)
	for i := 2; i <= 10; i++ {
		if num > max {
			max = num
		}
		fmt.Print(" Num ", i, ": ")
		fmt.Scan(&num)
	}

	fmt.Println("The largest value entered is:", max)

}
