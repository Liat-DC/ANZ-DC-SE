package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var text string = "What do you call a pony with a cough? A little horse."

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n') // Reads until newline
	word = strings.TrimSuffix(word, "\n")
	/*
		var word string
		fmt.Print("Enter a word: ")
		fmt.Scan(&word)
	*/

	if strings.Contains(text, word) {
		fmt.Println("The word appears in the text")
	} else {
		fmt.Println("The word doesn't appears in the text")

	}
}
