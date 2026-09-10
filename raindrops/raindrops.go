// package raindrops

// import "strconv"

// func Convert(number int) string {
// 	if number%3 == 0 && number%5 != 0 && number%7 != 0 {
// 		return "Pling"
// 	} else if number%3 == 0 && number%5 == 0 && number%7 == 0 {
// 		return "PlingPlangPlong"
// 	} else if number%3 == 0 && number%5 == 0 {
// 		return "PlingPlang"
// 	} else if number%3 == 0 && number%7 == 0 {
// 		return "PlingPlong"
// 	}
// 	if number%5 == 0 && number%7 != 0 && number%3 != 0 {
// 		return "Plang"
// 	} else if number%5 == 0 && number%7 == 0 {
// 		return "PlangPlong"
// 	}
// 	if number%7 == 0 {
// 		return "Plong"
// 	}
// 	return strconv.Itoa(number)

// }

package raindrops

import "strconv"

func Convert(number int) string {
	result := ""

	if number%3 == 0 {
		result += "Pling"
	}

	if number%5 == 0 {
		result += "Plang"
	}

	if number%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		return strconv.Itoa(number)
	}

	return result
}
