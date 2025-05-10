package problem4

import "math"

func Solve(input int) int {
	maxLoop := int(math.Pow10(input))
	maxPalindrome := 0
	for i := maxLoop - 1; i > 0; i-- {
		for j := i; j > 0; j-- {
			product := i * j
			if isPalindromeNumeric(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}
	return maxPalindrome
}

func isPalindromeNumeric(n int) bool {
	if n < 0 {
		return false
	}

	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n = n / 10
	}
	return original == reversed
}
