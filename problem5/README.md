# Problem 5: Smallest Multiple

## Problem Description
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to n?

## Algorithm
The solution uses the concept of least common multiple (LCM):

1. Initialize the LCM to 1
2. Iterate through numbers from 1 to n
3. For each number i, update the LCM to be the least common multiple of i and the current LCM
4. Return the final LCM

The least common multiple of two numbers a and b is calculated using the formula:
LCM(a, b) = (a * b) / GCD(a, b)

Where GCD is the greatest common divisor, calculated using the Euclidean algorithm.

## Implementation
The implementation leverages the LCM function from the common package, which:
1. Uses the GCD function to find the greatest common divisor
2. Calculates the LCM using the formula: LCM(a, b) = (a / GCD(a, b)) * b
3. Handles edge cases (when either number is 0)

This approach efficiently finds the smallest number divisible by all numbers from 1 to n by iteratively building up the LCM.