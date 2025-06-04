# Problem 7: 10001st Prime

## Problem Description
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the nth prime number?

## Algorithm
The solution uses a straightforward approach to find the nth prime number:

1. Start with count = 1 (accounting for 2 as the first prime)
2. Iterate through odd numbers starting from 3 (since even numbers greater than 2 are not prime)
3. For each number:
   - Check if it is prime using the IsPrime function
   - If it is prime, increment the count
   - If the count reaches n, return the current number as the answer
4. Return 0 if no prime is found (which should not happen for valid inputs)

## Implementation
The implementation optimizes the search by:
1. Only checking odd numbers (skipping even numbers which are never prime except for 2)
2. Using the efficient IsPrime function from the common package
3. Immediately returning when the nth prime is found

The IsPrime function itself is optimized by:
1. Handling edge cases (numbers less than 2)
2. Quickly identifying even numbers (except 2) as non-prime
3. Only checking odd divisors up to the square root of the number

This approach is efficient for finding prime numbers within the typical range of inputs for this problem.