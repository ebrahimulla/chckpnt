package piscine 


func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	len1, len2 := len(slice1), len(slice2)
	i, j := 0, 0

	// Determine which slice is larger
	startWithFirst := len1 >= len2

	// Alternate elements from both slices
	for i < len1 || j < len2 {
		if startWithFirst && i < len1 {
			result = append(result, slice1[i])
			i++
		} else if j < len2 {
			result = append(result, slice2[j])
			j++
		}

		if !startWithFirst && i < len1 {
			result = append(result, slice1[i])
			i++
		} else if j < len2 {
			result = append(result, slice2[j])
			j++
		}
	}

	return result
}