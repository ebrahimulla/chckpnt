package main

import (
	"os"
	"strconv"
    "fmt"
)

func main() {
	// Ensure there's exactly one argument
	if len(os.Args) != 2 {
		return
	}

	// Convert the argument to an integer
	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	// Find and print the prime factors
	primeFactors := getPrimeFactors(n)
	if len(primeFactors) == 0 {
		return
	}

	for i, factor := range primeFactors {
		if i > 0 {
			fmt.Print("*")
		}
		fmt.Print(factor)
	}
	fmt.Print("\n")
}

// Function to find prime factors
func getPrimeFactors(n int) []int {
	var factors []int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
