# Problem 9: Special Pythagorean Triplet

## Problem Description
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a² + b² = c²

For example, 3² + 4² = 9 + 16 = 25 = 5².

There exists exactly one Pythagorean triplet for which a + b + c = n.
Find the product abc.

## Algorithm
The solution uses a brute force approach with three nested loops:

1. Iterate through all possible values of a from 1 to n-2
2. For each a, iterate through all possible values of b from a+1 to n-1
3. For each a and b, calculate c = n - a - b
4. Check if the triplet (a, b, c) satisfies the Pythagorean condition: a² + b² = c²
5. When a valid triplet is found, return the product a * b * c

## Implementation
The implementation uses three nested loops to explore all possible combinations, with constraints to ensure:
1. a < b < c (by starting b at a+1 and c at b+1)
2. a + b + c = n (by calculating c as n - a - b)
3. a² + b² = c² (the Pythagorean condition)

This approach guarantees finding the unique triplet if it exists. While the brute force method has O(n³) time complexity in the worst case, it's efficient enough for the typical input values in this problem.