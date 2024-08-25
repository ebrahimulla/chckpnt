package piscine

func RepeatAlpha(s string) string {
	var result []rune

	for _, char := range s {
		var index int
		if char >= 'a' && char <= 'z' {
			index = int(char - 'a' + 1)
		} else if char >= 'A' && char <= 'Z' {
			index = int(char - 'A' + 1)
		} else {
			result = append(result, char)
			continue
		}

		for i := 0; i < index; i++ {
			result = append(result, char)
		}
	}

	return string(result)
}