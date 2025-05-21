package problem10

import "euler/common"

func Solve(input int) int {
	acc := 0
	for i := 1; i <= input; i++ {
		if common.IsPrime(i) {
			acc += i
		}
	}
	return acc
}
