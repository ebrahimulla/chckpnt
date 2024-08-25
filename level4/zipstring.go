package piscine

import (
    "strconv"
)

func ZipString(s string) string {
	result := ""
	n := len(s)
	i := 0

	for i < n {
		count := 1
		for i+1 < n && s[i] == s[i+1] {
			count++
			i++
		}
		result += strconv.Itoa(count) + string(s[i])
		i++
	}

	return result
}
