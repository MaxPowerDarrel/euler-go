package problem9

func Solve(input int) int {
	for i := 1; i <= input; i++ {
		for j := i + 1; j <= input; j++ {
			for k := j + 1; k <= input; k++ {
				if i+j+k == input && i*i+j*j == k*k {
					return i * j * k
				}
			}
		}
	}
	return 0
}
