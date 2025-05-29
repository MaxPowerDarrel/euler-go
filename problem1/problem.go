package problem1

import (
	calc "euler/common"
)

var divisors = []int{3, 5}

func Solve(input int) int {
	numbers := make([]int, input)
	for i := 1; i < input; i++ {
		if calc.IsDivisibleByAny(i, divisors) {
			numbers[i] = i
		}
	}
	return calc.Sum(numbers)
}
