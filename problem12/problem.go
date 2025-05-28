package problem12

import "math"

func Solve(input int) int {
	i := 1
	triangleNumber := 0
	for {
		if triangleNumber < 0 {
			panic("Negative triangle number")
		}
		triangleNumber += i
		if countDivisors(triangleNumber) > input {
			return triangleNumber
		}
		i++
	}
}

func countDivisors(n int) int {
	totalDivisors := 0
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			totalDivisors++
		}
	}
	return totalDivisors * 2
}
