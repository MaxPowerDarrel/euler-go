package problem19

func Solve() int {
	total, dayOfWeek := 0, 2
	for year := 1900; year < 2001; year++ {
		for month := 1; month <= 12; month++ {
			dayOfWeek = offSetDayOfWeek(month, year, dayOfWeek)
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

func offSetDayOfWeek(month, year, dayOfWeek int) int {
	dayOfWeek += monthToOffset(month, year)
	if dayOfWeek > 7 {
		dayOfWeek -= 7
	}
	return dayOfWeek
}

func monthToOffset(month, year int) int {
	switch month {
	case 4, 6, 9, 11:
		return 2
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return 1
		} else {
			return 0
		}
	default:
		return 3
	}
}
