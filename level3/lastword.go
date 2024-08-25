package piscine

func LastWord(s string) string {
	if s == "" {
		return "\n"
	}
	length := len(s)
	end := length - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return "\n"
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	lastWord := s[start+1 : end+1]
	return lastWord + "\n"
}