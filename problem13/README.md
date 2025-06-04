# Problem 13: Large Sum

## Problem Description
Work out the first ten digits of the sum of one-hundred 50-digit numbers.

## Algorithm
The solution uses a straightforward approach with big integers:

1. Store the 100 large numbers (each 50 digits) as strings
2. Convert each string to a big.Int using Go's math/big package
3. Sum all the big.Int numbers
4. Convert the final sum back to a string
5. Return the first 10 digits of the sum

## Implementation
The implementation leverages Go's `big.Int` type to handle the large numbers without precision loss. This is necessary because:
1. Standard integer types would overflow with 50-digit numbers
2. Floating-point types would lose precision
3. The `big.Int` package allows for arbitrary-precision arithmetic, ensuring accurate results

The solution efficiently handles the conversion between string representation and big integers, making it both accurate and performant.