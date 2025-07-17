package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n') // Reads until newline
	input = strings.TrimSuffix(input, "\n")

	//a.
	fmt.Println("Number of characters in the input:", len(input))

	//b.
	numNoSpaces := len(input)
	for _, r := range input {
		if r == ' ' {
			numNoSpaces--
		}
	}
	fmt.Println("Number of characters excluding spaces:", numNoSpaces)

	//c.
	numOfVowels := 0
	for _, r := range input {
		if r == 'a' || r == 'e' || r == 'i' || r == 'u' || r == 'o' {
			numOfVowels++
		}
	}
	fmt.Println("Number of vowels:", numOfVowels)

	//d.
	numOfWords := 0
	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			numOfWords++
			for input[i] == ' ' {
				i++
			}
		}
	}
	if len(input) > 0 {
		fmt.Println("Number of words:", numOfWords+1)
	} else {
		fmt.Println("Number of words:", 0)

	}
}
