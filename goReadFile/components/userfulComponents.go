package components

func GetNumberSuffix(n int) string {
    switch n {
    case 1:
        return "st"
    case 2:
        return "nd"
    case 3:
        return "rd"
    }
    return "th"
}