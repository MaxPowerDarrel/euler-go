package problem15

import calc "euler/common"

func Solve(input int) int64 {
	return calc.Combinations(input*2, input).Int64()
}
