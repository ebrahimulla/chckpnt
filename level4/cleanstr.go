package main

import (
	"os"

	"github.com/01-edu/z01"
)

// main function to process the input and format it with one space between words
func main() {
	if len(os.Args) != 2 {
		// If not exactly one argument, print newline and exit
		z01.PrintRune('\n')
		return
	}

	input := os.Args[1]
	n := len(input)

	// Skip leading spaces and tabs
	startIndex := 0
	for startIndex < n && (input[startIndex] == ' ' || input[startIndex] == '\t') {
		startIndex++
	}

	if startIndex == n {
		// If the string is only spaces or tabs
		z01.PrintRune('\n')
		return
	}

	inWord := false

	for i := startIndex; i < n; i++ {
		if input[i] != ' ' && input[i] != '\t' {
			if !inWord {
				// Starting a new word
				if i > startIndex && (input[i-1] == ' ' || input[i-1] == '\t') {
					// Print single space before the new word if it's not the first word
					z01.PrintRune(' ')
				}
				inWord = true
			}
			// Print the current character of the word
			z01.PrintRune(rune(input[i]))
		} else {
			// We encountered a space or tab
			inWord = false
		}
	}

	// Print newline at the end
	z01.PrintRune('\n')
}
