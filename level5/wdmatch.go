package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	target := os.Args[1]
	source := os.Args[2]

	if canForm(target, source) {
		printString(target)
	}
}

// canForm checks if the target string can be formed from the source string
// by respecting the order of characters.
func canForm(target, source string) bool {
	tIndex := 0
	sIndex := 0

	for tIndex < len(target) && sIndex < len(source) {
		if target[tIndex] == source[sIndex] {
			tIndex++
		}
		sIndex++
	}

	return tIndex == len(target)
}

// printString prints the given string followed by a newline
func printString(s string) {
	for _, ch := range s {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
