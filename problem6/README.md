# Problem 6: Sum Square Difference

## Problem Description
The sum of the squares of the first ten natural numbers is,
1² + 2² + ... + 10² = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)² = 55² = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first n natural numbers and the square of the sum.

## Algorithm
The solution uses a straightforward approach:

1. Initialize two variables:
   - `sum`: to store the sum of the first n natural numbers
   - `sumSquares`: to store the sum of the squares of the first n natural numbers
2. Iterate through numbers from 1 to n:
   - Add each number to `sum`
   - Add the square of each number to `sumSquares`
3. Calculate the square of the sum: `sum²`
4. Return the difference: `sum² - sumSquares`

## Implementation
The implementation is simple and efficient, with O(n) time complexity. For very large values of n, there are mathematical formulas that could be used instead:
- Sum of first n natural numbers: n(n+1)/2
- Sum of squares of first n natural numbers: n(n+1)(2n+1)/6

However, the current implementation is clear and efficient for the typical range of inputs in this problem.