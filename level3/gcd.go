package piscine

func Gcd(a, b uint) uint {
    if b == 0 {
        return a
    }
    return Gcd(b, a % b)
}

func Gcd(a, b uint) uint {
    // Handle the case where either number is zero
    if a == 0 || b == 0 {
        return 0
    }

    // Euclidean algorithm to find GCD
    for b != 0 {
        a, b = b, a % b
    }
    return a
}