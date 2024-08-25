package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, arg := range os.Args[1:] {
		wordStart := 0
		for i := 0; i <= len(arg); i++ {
			if i == len(arg) || arg[i] == ' ' {
				if wordStart < i {
					word := arg[wordStart:i]
					processWord(word)
				}
				if i < len(arg) {
					z01.PrintRune(' ')
				}
				wordStart = i + 1
			}
		}
		z01.PrintRune('\n')
	}
}

// processWord transforms a word to lowercase and capitalizes the last letter if it's a letter.
func processWord(word string) {
	length := len(word)
	if length == 0 {
		return
	}

	// Convert all characters to lowercase
	for i := 0; i < length-1; i++ {
		ch := word[i]
		if ch >= 'A' && ch <= 'Z' {
			ch = ch + 'a' - 'A'
		}
		z01.PrintRune(rune(ch))
	}

	// Capitalize the last letter if it's a letter
	lastChar := word[length-1]
	if lastChar >= 'a' && lastChar <= 'z' {
		lastChar = lastChar - ('a' - 'A')
	}
	z01.PrintRune(rune(lastChar))
}
