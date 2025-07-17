package main

import "fmt"

const SIZE int = 5

// The function takes an array of integers and
// returns the largest value
func find_max(arr [SIZE]int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

// The function searches for a value in an array and
// returns its index, or -1 if not found
func find(arr [SIZE]int, val int) int {
	for i, num := range arr {
		if num == val {
			return i
		}
	}
	return -1
}

// The function returns true if there are duplicate values in an array, or false otherwise
func duplicate_values(arr [SIZE]int) bool {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func main() {

	nums := [SIZE]int{}
	fmt.Println("Please enter", SIZE, "integers")
	// read values
	for i := range nums {
		fmt.Scan(&nums[i])
	}

	// print the entire array
	fmt.Print("The array is: ")
	fmt.Println(nums)

	// find and print the largest value
	largest := find_max(nums)
	fmt.Println("The largest number is:", largest)

	// search for a given number and its index
	fmt.Print("Enter a value to search: ")
	var val int
	fmt.Scan(&val)
	index := find(nums, val)
	if index != -1 {
		fmt.Println("The value", val, "was found in index", index)
	} else {
		fmt.Println("The value", val, "was not found in the array")
	}

	// check for duplicate values

	if duplicate_values(nums) {
		fmt.Println("There are duplicate values in the array")
	} else {
		fmt.Println("There are NO duplicate values in the array")
	}

}
