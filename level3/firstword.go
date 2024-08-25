package piscine

func FirstWord(s string) string {
	if s == "" {
		return "\n"
	}
	start := 0
	length := len(s)
	for start < length && s[start] == ' ' {
		start++
	}
	if start == length {
		return "\n"
	}
	end := start
	for end < length && s[end] != ' ' {
		end++
	}
	firstWord := s[start:end]
	return firstWord + "\n"
}