package piscine

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	start := 0
	for i := 0; i <= len(s); i++ {
		// If we encounter a space or reach the end of the string
		if i == len(s) || s[i] == ' ' {
			if start < i {
				word := s[start:i]
				if len(word) > 0 {
					firstChar := word[0]
					if !((firstChar >= 'A' && firstChar <= 'Z') || (firstChar < 'A' || firstChar > 'Z') && (firstChar < 'a' || firstChar > 'z')) {
						return false
					}
				}
			}
			start = i + 1
		}
	}
	return true
}