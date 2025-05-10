package problem2

func Solve(input int) int {
	c := generateFibonacciSequence(input)
	total := 0
	for n := range c {
		if n%2 == 0 {
			total += n
		}
	}
	return total
}

// generateFibonacciSequence returns a channel that produces Fibonacci numbers
// less than the specified limit. If maxNumber is <= 0, the channel is closed immediately.
func generateFibonacciSequence(maxNumber int) <-chan int {
	const (
		initialFirst  = 0
		initialSecond = 1
	)

	if maxNumber <= 0 {
		c := make(chan int)
		close(c)
		return c
	}

	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		first, second := initialFirst, initialSecond
		for sum := first + second; sum < maxNumber; sum = first + second {
			resultChan <- sum
			first, second = second, sum
		}
	}()

	return resultChan
}
