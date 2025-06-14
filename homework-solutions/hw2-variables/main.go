/*
This code includes a proposed solution for the four Questions of Exercise 2 - Variables
*/

package main

import "fmt"

func main() {

	///////////////////////////////////////////////////
	// Question 1
	var age = 15
	var height = 1.6

	fmt.Println(age)
	fmt.Println(height)

	///////////////////////////////////////////////////
	// Question 2
	var a, b int = 5, 8

	fmt.Println("a + b =", a+b)

	///////////////////////////////////////////////////
	// Question 3
	score := 10
	bonus := 5
	score += bonus // score = score + bonus
	fmt.Println("The score after bonus is:", score)

	///////////////////////////////////////////////////
	// Question 4 - Assigning a float to score

	/* To make score a float (e.g. 10.8), you can either:
	- explicitly declare it as float64, e.g. var score float64 = 10.8
	- OR use a float literal when using shorthand declaration (:=)
	*/

	score2 := 10.8 // score2 will be of type float64

	/* In Go, you cannot mix int and float64 types in arithmetic operations
	without explicit conversion. So, to add bonus to score2,
	bonus also needs to be a float64.
	*/

	bonus2 := 5.0 // Go infers the literal 5.0 as float64

	// Now that both score2 and bonus2 are float64, we can safely add them.
	score2 = score2 + bonus2
	fmt.Println("The score after bonus is:", score2)

	// Alternatively, you can keep bonus as an int and explicitly convert it
	// to float64 just in the expression.

	score3 := 10.8 // float64
	bonus3 := 5    // int
	score3 = score3 + float64(bonus3)
	fmt.Println("The score after bonus is:", score3)

}

