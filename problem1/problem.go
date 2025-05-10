package problem1

import (
	calc "euler/common"
)

var divisors = []int{3, 5}

func Solve(input int) int {
	var numbers []int
	for i := 1; i < input; i++ {
		if calc.IsDivisibleByAny(i, divisors) {
			numbers = append(numbers, i)
		}
	}
	return calc.Sum(numbers)
}
