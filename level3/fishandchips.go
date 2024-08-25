package piscine

func FishAndChips(n int) string {
    if n < 0{
        return ("error: number is negative")
    }

    var res string

    if n % 2 == 0{
        res += "fish"
    }
    if n % 3 == 0{
        if res != ""{
            res += " and "
        }
        res += "chips"
    }
    if res == ""{
        res += "error: non divisible"
    }
    return res
}