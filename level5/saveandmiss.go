package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	result := ""
	save := true

	for i := 0; i < len(arg); i += num {
		if save {
			end := i + num
			if end > len(arg) {
				end = len(arg)
			}
			result += arg[i:end]
		}
		save = !save
	}

	return result
}
