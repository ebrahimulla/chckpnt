package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	
	// Handle negative numbers
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	// Convert integer to string
	result := ""
	for n > 0 {
		digit := n % 10
		result = string(digit+'0') + result
		n = n / 10
	}

	// Prepend the sign if necessary
	return sign + result
}