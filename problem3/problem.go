package problem3

import (
	"euler/common"
	"math"
)

func Solve(input int) int {
	targetNumber := input
	targetNumberSqrt := int(math.Sqrt(float64(targetNumber)))
	var highestPrimeFactor int
	for i := 1; i < targetNumberSqrt; i++ {
		if targetNumber%i == 0 && common.IsPrime(i) && i > highestPrimeFactor {
			highestPrimeFactor = i
		}
	}
	return highestPrimeFactor
}
