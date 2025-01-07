package ordinal

import "strconv"

func intToOrdinal(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%10 != 11 {
			suffix = "st"
		}
	case 2:
		if x%10 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%10 != 13 {
			suffix = "rd"
		}
	}
	return strconv.Itoa(x) + suffix
}
