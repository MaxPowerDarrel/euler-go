package common

import (
	"math"
	"math/big"
)

func IsDivisibleByAny(n int, divisors []int) bool {
	for _, divisor := range divisors {
		if n%divisor == 0 {
			return true
		}
	}
	return false
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a / GCD(a, b)) * b
}

func Combinations(n, r int) *big.Int {
	// Calculate C(n,r) = n! / (r! * (n-r)!) more efficiently
	// Instead of computing full factorials, we can simplify:
	// C(n,r) = (n * (n-1) * ... * (n-r+1)) / (r * (r-1) * ... * 1)
	result := big.NewInt(1)
	for i := 0; i < r; i++ {
		result = result.Mul(result, big.NewInt(int64(n-i)))
		result = result.Div(result, big.NewInt(int64(i+1)))
	}

	return result
}
