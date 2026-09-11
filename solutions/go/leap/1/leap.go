package leap

// IsLeapYear calculate what year es divisble for 4 and 400
func IsLeapYear(year int) bool {
    if year % 400 == 0{
        return true
    }
    if year % 100 == 0{
        return false
    }
    if year % 4 == 0{
        return true
    }
    return false
}
