package problem6

func Solve(input int) int {
	sum := 0
	sumSquares := 0
	for i := 1; i <= input; i++ {
		sum += i
		sumSquares += i * i
	}
	return sum*sum - sumSquares
}
