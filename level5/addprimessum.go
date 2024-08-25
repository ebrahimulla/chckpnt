package main

import (
	"os"
	"github.com/01-edu/z01"
)

// isPrime checks if a number is a prime number
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

// convertIntToRuneSlice converts an integer to a slice of runes
func convertIntToRuneSlice(num int) []rune {
	if num == 0 {
		return []rune{'0'}
	}
	var runes []rune
	if num < 0 {
		num = -num
	}
	for num > 0 {
		runes = append([]rune{rune(num%10 + '0')}, runes...)
		num /= 10
	}
	return runes
}

// printNumber prints a number using z01.PrintRune
func printNumber(num int) {
	runes := convertIntToRuneSlice(num)
	for _, r := range runes {
		z01.PrintRune(r)
	}
}

func main() {
	if len(os.Args) != 2 {
		printNumber(0)
		z01.PrintRune('\n')
		return
	}

	arg := os.Args[1]
	num := 0
	for _, r := range arg {
		if r < '0' || r > '9' {
			printNumber(0)
			z01.PrintRune('\n')
			return
		}
		num = num*10 + int(r-'0')
	}
	if num <= 0 {
		printNumber(0)
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printNumber(sum)
	z01.PrintRune('\n')
}
