# Problem 14: Longest Collatz Sequence

## Problem Description
The Collatz sequence is defined for any positive integer n:
- If n is even, the next term is n/2
- If n is odd, the next term is 3n+1

The sequence continues until it reaches 1. The problem asks which starting number under one million produces the longest chain.

## Algorithm
The solution uses a brute force approach with concurrency:

1. For each number from 1 to 1,000,000:
   - Create a goroutine to calculate the length of its Collatz sequence
   - Each goroutine follows the Collatz rules until reaching 1, counting the steps
   - The result (starting number and sequence length) is sent to a channel
2. The main function collects all results from the channel
3. It keeps track of the maximum sequence length and the corresponding starting number
4. Finally, it returns the number with the longest sequence

## Implementation
The implementation leverages Go's concurrency model with goroutines and channels to process all numbers in parallel, making efficient use of multiple CPU cores. This significantly speeds up the computation compared to a sequential approach.