// This program includes code demonstrate in class on 24 June 2025 (week 6)

/* String example - multiline */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Example:
	// -Defining a string variable
	// -Assigning a new value to a string
	// -Printing a single character

	str1 := "Hello, World!" // Regular string
	str2 := `Multiline
string example.` // Raw string

	fmt.Println(str1)
	fmt.Println(str2)

	var s string = "This is a string"
	fmt.Println(s)
	s = "test"
	fmt.Println(s)
	fmt.Println(s[0])        // Will print the character code, i.e.: 116
	fmt.Printf("%c\n", s[0]) // Will print the actual letter, i.e.: t

	/********************************************************************************/
  // Example: iterating through strings

	for i := range len(s) { // Iterating using an INDEX - will iterate over BYTES
		fmt.Print(i, ": ")
		fmt.Printf("%c\n", s[i])
	}

	s = "élite"
	for _, r := range s { // Iterating through RUNES
		fmt.Printf("Letter: %c\n", r)
	}

	/********************************************************************************/
  // Example: reading a full sentence from the user (string may include spaces)
  
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	text, _ := reader.ReadString('\n') // Reads until newline
	fmt.Println("You entered:", text)

}
