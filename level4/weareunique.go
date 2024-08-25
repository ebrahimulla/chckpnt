package piscine

import "strings"

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	uniqueChars := make(map[rune]bool)
	
	// Add characters from str1 that are not in str2
	for _, char := range str1 {
		if !strings.Contains(str2, string(char)) {
			uniqueChars[char] = true
		}
	}
	
	// Add characters from str2 that are not in str1
	for _, char := range str2 {
		if !strings.Contains(str1, string(char)) {
			uniqueChars[char] = true
		}
	}

	// Count unique characters
	count := len(uniqueChars)
	
	return count
}