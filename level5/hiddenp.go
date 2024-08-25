package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	if isHidden(s1, s2) {
		printResult(1)
	} else {
		printResult(0)
	}
}

// isHidden checks if s1 is hidden in s2
func isHidden(s1, s2 string) bool {
	s1Index := 0
	s2Index := 0

	// Iterate through s2 and try to match s1
	for s1Index < len(s1) && s2Index < len(s2) {
		if s1[s1Index] == s2[s2Index] {
			s1Index++
		}
		s2Index++
	}

	return s1Index == len(s1)
}

// printResult prints the result followed by a newline
func printResult(result int) {
	z01.PrintRune(rune(result + '0'))
	z01.PrintRune('\n')
}
