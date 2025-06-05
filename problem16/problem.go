package problem16

import (
	"euler/common"
	"math/big"
)

func Solve(input int64) int {
	bigNumber := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(input), nil).String()
	total := 0
	for _, char := range bigNumber {
		total += common.NumberCharToInt(char)
	}
	return total
}
