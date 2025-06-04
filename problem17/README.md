# Problem 17: Number Letter Counts

## Problem Description
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters.

## Algorithm
The solution uses a straightforward approach:

1. Define constants for the word representation of each number (ONE, TWO, THREE, etc.)
2. Implement a function `numberToString` that converts a number to its word representation:
   - Handle special case for 1000 ("one thousand")
   - For other numbers, break them down into hundreds, tens, and ones
   - For hundreds, add the word representation of the digit followed by "hundred"
   - If there's a remainder after hundreds, add "and"
   - For tens and ones:
     - If the remainder is 20 or greater, handle tens and ones separately
     - If the remainder is less than 20, handle it directly using predefined constants
3. Implement a helper function `grabStringPart` that returns the word representation for a given number
4. The main `Solve` function:
   - Iterates through all numbers from 1 to the input value
   - Converts each number to its word representation using `numberToString`
   - Calculates the length of each word representation
   - Sums up all the lengths

The solution has a time complexity of O(n) where n is the input value, as it processes each number from 1 to n exactly once.