// This program generates a random number between 1 and 10 (inclusive), then give the user three attempts to guess the number.
// NOTE: this program doesn't use a for loop. Using a loop will make a better solution for this type of problems.

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Generate a random number between 1 and 10
	secret := rand.Intn(10) + 1

	var guess int

	// First attempt
	fmt.Print("Attempt 1: Enter your guess (1-10): ")
	fmt.Scan(&guess)
	if guess == secret {
		fmt.Println("Correct! You guessed the number.")
		return
	} else if guess < secret {
		fmt.Println("The correct number is greater.")
	} else {
		fmt.Println("The correct number is smaller.")
	}

	// Second attempt
	fmt.Print("Attempt 2: Enter your guess (1-10): ")
	fmt.Scan(&guess)
	if guess == secret {
		fmt.Println("Correct! You guessed the number.")
		return
	} else if guess < secret {
		fmt.Println("The correct number is greater.")
	} else {
		fmt.Println("The correct number is smaller.")
	}

	// Third attempt
	fmt.Print("Attempt 3: Enter your guess (1-10): ")
	fmt.Scan(&guess)
	if guess == secret {
		fmt.Println("Correct! You guessed the number.")
		return
	} else {
		fmt.Printf("Sorry, you're out of tries. The correct number was %d.\n", secret)
	}

}
