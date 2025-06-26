// The program prints the sum of numbers from 1 to a given number (inclusive).
// For example, if the user entered 7 the output should be 28 (1 + 2 + 3 + 4 + 5 + 6 + 7)


package main

import (
	"fmt"
)

func main() {

		var num int

		fmt.Println("Please enter a positive integer")
		fmt.Scan(&num)

		var sum int = 0

		for i := 1; i <= num; i++ {
			sum = sum + i
		}

		fmt.Println("The sum is", sum)

}
