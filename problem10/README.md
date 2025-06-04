# Problem 10: Summation of Primes

## Problem Description
Find the sum of all the prime numbers below n.

## Algorithm
The solution uses a straightforward approach:

1. Iterate through all numbers from 1 to n
2. For each number, check if it is prime
3. If the number is prime, add it to a running sum
4. Return the final sum

## Implementation
The implementation leverages the `IsPrime` function from the common package, which efficiently determines if a number is prime by:
1. Handling edge cases (numbers less than 2)
2. Quickly identifying even numbers (except 2) as non-prime
3. Only checking odd divisors up to the square root of the number

This approach is efficient for the range of numbers typically used in this problem. For extremely large inputs, a more optimized approach like the Sieve of Eratosthenes might be more efficient, but the current implementation provides a good balance of simplicity and performance.