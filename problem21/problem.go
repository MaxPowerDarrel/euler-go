package problem21

import "euler/common"

// Solve finds the sum of all amicable numbers under 10,000.
// Two numbers are amicable if the sum of the divisors of each
// is equal to the other number.
func Solve() int {
	const limit = 10000

	// Pre-compute sums of proper divisors for all numbers
	divisorSums := make(map[int]int, limit)

	var total int
	for i := 1; i < limit; i++ {
		sumA := getSumOfDivisors(i, divisorSums)
		sumB := getSumOfDivisors(sumA, divisorSums)

		// Check if they form an amicable pair
		if sumA != i && sumB == i {
			total += i
		}
	}

	return total
}

// getSumOfDivisors returns the sum of proper divisors for n,
// using memoization to avoid recalculation.
func getSumOfDivisors(n int, cache map[int]int) int {
	if n <= 0 {
		return 0
	}

	if sum, exists := cache[n]; exists {
		return sum
	}

	sum := common.Sum(common.FindDivisors(n))
	cache[n] = sum
	return sum
}
