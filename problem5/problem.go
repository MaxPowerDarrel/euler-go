package problem5

import "euler/common"

func Solve(input int) int {
	lcm := 1
	for i := 1; i <= input; i++ {
		lcm = common.LCM(i, lcm)
	}
	return lcm
}
