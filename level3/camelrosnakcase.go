package piscine

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return s
	}

	// Initialize an empty slice of bytes to build the result
	var result []byte

	// Iterate over each character in the input string
	for i := 0; i < len(s); i++ {
		c := s[i]

		// Check if the character is not a letter
		if (c < 'A' || c > 'Z') && (c < 'a' || c > 'z') {
			return s
		}

		// Check if the character is uppercase
		if c >= 'A' && c <= 'Z' {
			// If it's the first character, keep it as is
			if i == 0 {
				result = append(result, c)
			} else {
				// If the previous character is also uppercase, it's not camelCase
				if i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z' {
					return s
				}
				if i == len(s)-1 && s[i] >= 'A' && s[i] <= 'Z' {
					return s
				}
				// Append an underscore followed by the uppercase character
				result = append(result, '_', c)
			}
		} else {
			// Append lowercase characters as is
			result = append(result, c)
		}
	}

	// Convert the byte slice back to a string and return it
	return string(result)
}
