package piscine

func RetainFirstHalf(str string) string {
    half := len(str)
    if half == 0{
        return ""
    }
    if half == 1{
        return str
    }
    cal := half / 2
    return str[:cal]
}