package piscine

func DigitLen(n, base int) int {
    count := 0
    if n < 0{
        n = -n
    }
    if base < 2 || base > 36 {
		return -1
    }else{
		for n > 0{
            n /= base
            count++
        }
        return count
    }
}
