package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    // Check if the number of arguments is exactly 2
    if len(os.Args) != 3 {
        return
    }

    str1 := os.Args[1]
    str2 := os.Args[2]
    
    // Use a map to keep track of characters already printed
    printed := make(map[rune]bool)
    
    for _, ch := range str1 {
        // Check if the character is in str2 and has not been printed yet
        if contains(str2, ch) && !printed[ch] {
            z01.PrintRune(ch)  // Print the character
            printed[ch] = true // Mark the character as printed
        }
    }
    
    z01.PrintRune('\n') // Print newline at the end
}

// contains checks if a rune is present in a string
func contains(s string, ch rune) bool {
    for _, c := range s {
        if c == ch {
            return true
        }
    }
    return false
}
