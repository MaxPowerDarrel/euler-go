package problem14

func Solve() (int, int) {
	const iterations = 1_000_000
	channel := make(chan [2]int, iterations)
	for i := 1; i <= iterations; i++ {
		go collatzSequence(i, channel)
	}
	maxIterations := 1
	maxNumber := 0
	for i := 0; i < iterations; i++ {
		n := <-channel
		if n[1] > maxIterations {
			maxIterations = n[1]
			maxNumber = n[0]
		}

	}
	return maxNumber, maxIterations
}

func collatzSequence(n int, channel chan [2]int) {
	count := 1
	number := n
	for number != 1 {
		if number%2 == 0 {
			number /= 2
			count++
		} else {
			number = 3*number + 1
			count++
		}
	}
	channel <- [2]int{n, count}
}
