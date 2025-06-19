package main

import (
	"fmt"
)

func main() {

	var num int

	fmt.Println("Please enter a number")
	fmt.Scan(&num)
	if num < 0 {
		fmt.Println("The absolute value of ", num, "is", num*(-1))
	} else {
		fmt.Println("The absolute value of ", num, "is", num)
	}
}
