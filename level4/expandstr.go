package main

import (
	"os"

	"github.com/01-edu/z01"
)

// main function to process the input and display it with three spaces between words
func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	var words []string
	var word string

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(input[i])
		}
	}

	if len(word) > 0 {
		words = append(words, word)
	}

	if len(words) == 0 {
		return
	}

	for i, word := range words {
		if i > 0 {
			// Print three spaces between words
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
		// Print each character of the word
		for _, char := range word {
			z01.PrintRune(char)
		}
	}

	// Print newline at the end
	z01.PrintRune('\n')
}
