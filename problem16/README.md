# Problem 16: Power Digit Sum

## Problem Description
Calculate the sum of the digits of 2^n.

## Algorithm
The solution uses the following approach:

1. Calculate 2^n using Go's `math/big` package to handle large numbers
2. Convert the result to a string representation
3. Iterate through each character (digit) in the string
4. Convert each character to its numeric value by subtracting the ASCII value of '0'
5. Add each digit to a running total
6. Return the final sum

## Implementation
The implementation uses Go's `big.Int` type to handle the potentially very large number that results from calculating 2^n, especially for large values of n. This avoids any overflow issues that would occur with standard integer types.