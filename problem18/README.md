# Problem 18: Maximum Path Sum I

## Problem Description
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

```
   3
  7 4
 2 4 6
8 5 9 3
```

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom of the triangle.

## Algorithm
The solution uses a bottom-up dynamic programming approach:

1. Start with the triangle represented as a 2D array
2. Work from the bottom row upward
3. For each element in the current row, add the maximum of the two adjacent elements from the row below
4. Continue this process until reaching the top of the triangle
5. The single value at the top will be the maximum path sum

Specifically, the algorithm:
1. Iteratively reduces the triangle by combining rows from bottom to top
2. For each pair of adjacent rows:
   - Create a new row where each element is the sum of the element in the upper row plus the maximum of the two adjacent elements in the lower row
   - Replace the upper row with this new row
   - Remove the lower row
3. Continue until only one row (with one element) remains
4. Return this element as the maximum path sum

This approach has O(n²) time complexity, where n is the number of rows in the triangle, as we process each element in the triangle exactly once.