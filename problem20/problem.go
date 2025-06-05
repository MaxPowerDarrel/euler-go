package problem20

import "euler/common"

func Solve(input int) int {
	total := 0
	factorial := common.Factorial(input).String()
	for _, char := range factorial {
		total += common.NumberCharToInt(char)
	}
	return total
}
