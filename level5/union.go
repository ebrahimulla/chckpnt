package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    // Check if the number of arguments is exactly 2
    if len(os.Args) != 3 {
        z01.PrintRune('\n') // Print newline if the number of arguments is not 2
        return
    }

    str1 := os.Args[1]
    str2 := os.Args[2]

    // Use a map to keep track of characters already printed
    seen := make(map[rune]bool)
    
    // Function to process a string and print characters in order without duplicates
    processString := func(s string) {
        for _, ch := range s {
            if !seen[ch] {
                seen[ch] = true
                z01.PrintRune(ch)
            }
        }
    }

    // Process both strings
    processString(str1)
    processString(str2)
    
    z01.PrintRune('\n') // Print newline at the end
}
