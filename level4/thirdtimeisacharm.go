package piscine

import (
	"strings"
)


func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	var result strings.Builder

	for i := 2; i < len(str); i += 3 {
		result.WriteByte(str[i])
	}

	result.WriteByte('\n')
	return result.String()
}


func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	result := make([]rune, 0, len(str)/3)

	for i := 2; i < len(str); i += 3 {
		result = append(result, rune(str[i]))
	}

	return string(result) + "\n"
}