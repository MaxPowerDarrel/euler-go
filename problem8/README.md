# Problem 8: Largest Product in a Series

## Problem Description
The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

Find the n adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

## Algorithm
The solution uses a sliding window approach:

1. Store the 1000-digit number as a string
2. Iterate through all possible starting positions for n consecutive digits:
   - For each position, calculate the product of the n consecutive digits
   - Convert each character to its numeric value by subtracting the ASCII value of '0'
   - Multiply all n digits together
3. Keep track of the maximum product found
4. Return the maximum product

## Implementation
The implementation efficiently handles this problem by:
1. Using a single pass through the number with a sliding window of size n
2. Calculating the product for each window
3. Updating the maximum product when a larger one is found

This approach has O(m) time complexity, where m is the length of the number (1000 in this case), making it very efficient. The space complexity is O(1) as we only store a few variables regardless of the input size.