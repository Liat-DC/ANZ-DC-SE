// This program reads 5 numbers from the user and prints the largest and second largest numbers

package main

import (
	"fmt"
)

func main() {

	var num, max, second_max int

	var first_update bool = true  // a boolean variables that will be used as a flag to know if this is the first update of the second largest value. We need this flag because updating the largest value the first time is different from the second time forward

	// Readinf the first number and updating max
  fmt.Println("Please enter 5 numbers")
	fmt.Print(" Num 1: ")
	fmt.Scan(&num)
	max = num
  second_max = num
	for i := 2; i <= 5; i++ {
		fmt.Print(" Num ", i, ": ")
		fmt.Scan(&num)
		if num > max { // a new max was found: update both max and second_max
			second_max = max          
			first_update = false
			max = num
		} else if num < max && first_update { // first time found a number smaller than max: update second_max
			second_max = num
			first_update = false
		}
	}

	fmt.Println("The largest value entered is:", max)
	fmt.Println("The second largest value entered is:", second_max)

}
