package piscine

import (
    "strconv"
)

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	var result string

    if from == 1 && to == 1 {
        result += "1"
    

	} else if from <= to {
		for i := from; i <= to; i++ {
            
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i < to {
				result += ", "
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i > to {
				result += ", "
			}
		}
	}

	result += "\n"
	return result
}