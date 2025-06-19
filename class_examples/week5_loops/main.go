// This program includes code demonstrated in class on 17 June 2025

package main

import "fmt"

func main() {

	// Print the numbers from 0 to 9
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
	/*********************************************************************/

	// Print the numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
	/*********************************************************************/

	// Print the numbers from 10 to 1
	for i := 10; i >= 1; i-- {
		fmt.Print(i, " ")
	}

	fmt.Println()
	/*********************************************************************/

	// Print just the even numbers between 1 and 20
	for i := 2; i <= 20; i += 2 {
		fmt.Print(i, " ")
	}

	fmt.Println()
	/*********************************************************************/

	// Read 5 numbers from the user and print their sum and average

	var sum float64 = 0
	var avg float64 = 0
	var number float64

	fmt.Println("Enter 5 numbers")
	for i := 0; i < 5; i++ {
		fmt.Scan(&number)
		sum = sum + number
	}

	// calculate average
	avg = sum / 5
	// print sum
	fmt.Println("sum = ", sum)
	// print avg
	fmt.Println("average = ", avg)

	fmt.Println()
	/*********************************************************************/

	// Read positive numbers from the user until the number entered is -1 and print their sum and average

	var sum2, avg2 float64
	var number2 float64
	var counter int = 0

	fmt.Println("Please enter positive numbers. If you are done enter -1")

	for {
		// read number
		fmt.Scan(&number2)
		// if the number is -1 finish the loop
		if number2 == -1 {
			break
		}
		// update sum with the number
		sum2 = sum2 + number2
		counter = counter + 1 // counter++
	}

	avg2 = sum2 / float64(counter)

	// print sum
	fmt.Println("sum = ", sum2)
	// print avg
	fmt.Println("avg = ", avg2)

	fmt.Println()

	/*********************************************************************/

	//Example: nested loops - rectangle of 5x8 asteriks

	for i := 0; i < 5; i++ { // number of rows
		// print a line of 8 asteriks
		for j := 0; j < 8; j++ { // number of columns
			fmt.Print("* ")
		}
		// go the the next line
		fmt.Println()

	}

}
