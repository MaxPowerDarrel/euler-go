package problem17

const (
	ONE       = "ONE"
	TWO       = "TWO"
	THREE     = "THREE"
	FOUR      = "FOUR"
	FIVE      = "FIVE"
	SIX       = "SIX"
	SEVEN     = "SEVEN"
	EIGHT     = "EIGHT"
	NINE      = "NINE"
	TEN       = "TEN"
	ELEVEN    = "ELEVEN"
	TWELVE    = "TWELVE"
	THIRTEEN  = "THIRTEEN"
	FOURTEEN  = "FOURTEEN"
	FIFTEEN   = "FIFTEEN"
	SIXTEEN   = "SIXTEEN"
	SEVENTEEN = "SEVENTEEN"
	EIGHTEEN  = "EIGHTEEN"
	NINETEEN  = "NINETEEN"
	TWENTY    = "TWENTY"
	THIRTY    = "THIRTY"
	FORTY     = "FORTY"
	FIFTY     = "FIFTY"
	SIXTY     = "SIXTY"
	SEVENTY   = "SEVENTY"
	EIGHTY    = "EIGHTY"
	NINETY    = "NINETY"
	HUNDRED   = "HUNDRED"
	THOUSAND  = "THOUSAND"
	AND       = "AND"
)

func Solve(input int) int {
	totalLength := 0
	for i := 1; i <= input; i++ {
		totalLength += len(numberToString(i))
	}
	return totalLength
}

func numberToString(num int) string {
	if num == 0 {
		return ""
	}

	if num == 1000 {
		return ONE + THOUSAND
	}

	var parts []string

	// Handle hundreds
	hundreds, remainder := num/100, num%100
	if hundreds > 0 {
		parts = append(parts, grabStringPart(hundreds)+HUNDRED)
		if remainder > 0 {
			parts = append(parts, AND)
		}
	}

	// Handle tens and ones
	if remainder >= 20 {
		tens, ones := remainder/10*10, remainder%10
		parts = append(parts, grabStringPart(tens))
		if ones > 0 {
			parts = append(parts, grabStringPart(ones))
		}
	} else if remainder > 0 {
		// Handle 1-19 directly
		parts = append(parts, grabStringPart(remainder))
	}

	result := ""
	for _, part := range parts {
		result += part
	}
	return result
}

func grabStringPart(number int) string {
	switch number {
	case 1:
		return ONE
	case 2:
		return TWO
	case 3:
		return THREE
	case 4:
		return FOUR
	case 5:
		return FIVE
	case 6:
		return SIX
	case 7:
		return SEVEN
	case 8:
		return EIGHT
	case 9:
		return NINE
	case 10:
		return TEN
	case 11:
		return ELEVEN
	case 12:
		return TWELVE
	case 13:
		return THIRTEEN
	case 14:
		return FOURTEEN
	case 15:
		return FIFTEEN
	case 16:
		return SIXTEEN
	case 17:
		return SEVENTEEN
	case 18:
		return EIGHTEEN
	case 19:
		return NINETEEN
	case 20:
		return TWENTY
	case 30:
		return THIRTY
	case 40:
		return FORTY
	case 50:
		return FIFTY
	case 60:
		return SIXTY
	case 70:
		return SEVENTY
	case 80:
		return EIGHTY
	case 90:
		return NINETY
	case 100:
		return HUNDRED
	default:
		return ""
	}
}
