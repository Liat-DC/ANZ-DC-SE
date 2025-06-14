/*
This code includes a proposed solution for the three Questions of Exercise 1 - Basic Program
*/

package main

import "fmt"

func main() {

	// Question 1 - option 1
	fmt.Println("1")
	fmt.Println("22")
	fmt.Println("333")
	fmt.Println("4444")
	fmt.Println("55555")

	fmt.Println() // Prints a blank line (newline)

	// Question 1 - option 2
	fmt.Println("1\n22\n333\n4444\n55555")

	fmt.Println() // Prints a blank line (newline)

	// Question 2 - option 1 - hardcoded values
	fmt.Println("5 x 1 = 5")
	fmt.Println("5 x 2 = 10")
	fmt.Println("5 x 3 = 15")
	fmt.Println("5 x 4 = 20")
	fmt.Println("5 x 5 = 25")
	fmt.Println("5 x 6 = 30")
	fmt.Println("5 x 7 = 35")
	fmt.Println("5 x 8 = 40")
	fmt.Println("5 x 9 = 45")
	fmt.Println("5 x 10 = 50")

	fmt.Println() // Prints a blank line (newline)

	// Question 2 - option 2 - using the multiplication operator

	fmt.Println("5 x 1 =", 5*1)
	fmt.Println("5 x 2 =", 5*2)
	fmt.Println("5 x 3 =", 5*3)
	fmt.Println("5 x 4 =", 5*4)
	fmt.Println("5 x 5 =", 5*5)
	fmt.Println("5 x 6 =", 5*6)
	fmt.Println("5 x 7 =", 5*7)
	fmt.Println("5 x 8 =", 5*8)
	fmt.Println("5 x 9 =", 5*9)
	fmt.Println("5 x 10 =", 5*10)

	fmt.Println() // Prints a blank line (newline)

	// Question 3 - option 1 - using back ticks
	fmt.Println(`She said, "Hello!\nHow are you?"`)

	fmt.Println() // Prints a blank line (newline)

	// Quesion 3 - option 2 - using escape characters
	fmt.Println("She said, \"Hello!\\nHow are you?\"")

}

