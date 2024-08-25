package piscine

func HashCode(dec string) string {
	length := len(dec)
	var hashedString string

	for _, char := range dec {
		ascii := int(char)
		hashed := (ascii + length) % 127

		// Ensure the result is printable
		if hashed < 33 {
			hashed += 33
		}

		hashedString += string(rune(hashed))
	}

	return hashedString
}