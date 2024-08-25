package picine


func CheckNumber(arg string) bool {
    for _, x := range arg {
	    if x < 'a' || x > 'z' {
			if x == ' '{
            return false
            }
            return true
		}
	}
	return false
}