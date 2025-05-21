package problem7

import "euler/common"

func Solve(input int) int {
	count := 1
	for i := 3; count < input; i = i + 2 {
		if common.IsPrime(i) {
			count++
		}
		if count == input {
			return i
		}
	}
	return 0
}
