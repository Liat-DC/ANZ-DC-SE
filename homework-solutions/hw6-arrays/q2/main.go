package main

import "fmt"

func main() {
	var lottery = [6]int{5, 12, 23, 34, 45, 49}

	nums := [6]int{}
	fmt.Print("Enter your numbers: ")
	// read numbers
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	var count int = 0
	for _, win := range lottery {
		for _, num := range nums {
			if num == win {
				count++
			}
		}
	}
	fmt.Println("The winning numbers are: ", lottery)
	fmt.Println("You have", count, "matches")
}
