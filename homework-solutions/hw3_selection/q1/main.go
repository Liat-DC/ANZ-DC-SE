package main

import (
	"fmt"
)

func main() {

	var grade int

	fmt.Println("Please enter your grade")
	fmt.Scan(&grade)

	if grade >= 0 && grade <= 100 {
		if grade >= 85 {
			fmt.Println("Excellent work!")
		} else if grade >= 70 {
			fmt.Println("Good job!")
		} else if grade >= 50 {
			fmt.Println("You passed. Keep improving!")
		} else {
			fmt.Println("Don't give up. Keep trying!")
		}
		// Print sum of digits
		var sum int = 0
		if grade == 100 {
			sum = 1
		} else {
			sum = grade%10 + grade/10
		}
		fmt.Println("The sum of your grade's digits is", sum)
	} else {
		fmt.Println("Sorry. The grade is out of range.")
	}
}
