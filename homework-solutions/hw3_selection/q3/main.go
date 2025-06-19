// The program asks the user to enter three integers and
// prints the numbers in order from largest to smallest, labelling each as max, mid, and min.

package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var max, mid, min int

	fmt.Println("Please enter three integers")
	fmt.Scan(&a, &b, &c)

	if a > b && a > c {
		max = a
		if b > c {
			mid = b
			min = c
		} else {
			mid = c
			min = b
		}
	} else if b > a && b > c {
		max = b
		if a > c {
			mid = a
			min = c
		} else {
			mid = c
			min = a
		}
	} else {
		max = c
		if b > a {
			mid = b
			min = a
		} else {
			mid = a
			min = b
		}
	}

	fmt.Printf("max: %d\nmid: %d\nmin: %d\n", max, mid, min)
}
