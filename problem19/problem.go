package problem19

func Solve() int {
	total, dayOfWeek := 0, 2
	for year := 1900; year < 2001; year++ {
		for month := 1; month <= 12; month++ {
			dayOfWeek = offsetDayOfWeek(month, year, dayOfWeek)
			if year < 1901 {
				continue
			}
			if dayOfWeek == 1 {
				total++
			}
		}
	}
	return total
}

// offsetDayOfWeek calculates the day of the week for the first day of the given month
func offsetDayOfWeek(month, year, currentDayOfWeek int) int {
	offset := monthOffset(month, year)
	newDayOfWeek := currentDayOfWeek + offset

	// Normalize to 1-7 by looping around the start of the week again
	if newDayOfWeek > 7 {
		newDayOfWeek -= 7
	}

	return newDayOfWeek
}

// monthOffset returns the number of days to add based on the previous month's length
func monthOffset(month, year int) int {
	switch month {
	case 4, 6, 9, 11: // April, June, September, November (30 days)
		return 2 // 30 % 7 = 2
	case 2: // February
		if isLeapYear(year) {
			return 1 // 29 % 7 = 1
		}
		return 0 // 28 % 7 = 0
	default: // January, March, May, July, August, October, December (31 days)
		return 3 // 31 % 7 = 3
	}
}

// isLeapYear determines if a given year is a leap year
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
