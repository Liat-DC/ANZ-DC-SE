package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var letters = [26]int{} // letter counter

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n') // Reads until newline
	input = strings.ToUpper(input)

	for _, char := range input {
		if unicode.IsLetter(rune(char)) {
			letters[char-'A']++
		}
	}

	for i, c := range letters {
		if c > 0 {
			fmt.Printf("%c: ", i+'A')
			for j := 0; j < c; j++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}
}
