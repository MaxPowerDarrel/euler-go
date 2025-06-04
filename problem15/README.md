# Problem 15: Lattice Paths

## Problem Description
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner. How many such routes are there through a n×n grid?

## Algorithm
The solution uses a mathematical approach based on combinatorics:

1. In an n×n grid, to reach from the top-left to the bottom-right corner, you must make exactly 2n moves:
   - n moves to the right
   - n moves down
2. The order of these moves determines the unique path
3. This is equivalent to choosing which of the 2n positions will be "right" moves (or equivalently, which will be "down" moves)
4. Therefore, the number of paths is the binomial coefficient C(2n, n) = (2n)! / (n! * n!)

## Implementation
The implementation uses the `Combinations` function from the common package, which efficiently calculates the binomial coefficient C(2n, n) without computing full factorials. This avoids overflow issues that would occur with a naive implementation for large values of n.