# Problem 67: Maximum Path Sum II

## Problem Description
This problem is a larger version of Problem 18. By starting at the top of the triangle below and moving to adjacent numbers on the row below, find the maximum total from top to bottom.

The triangle contains 100 rows, making a brute-force approach (testing all routes) impossible as there are 2^99 possible routes.

## Algorithm
The solution uses the same bottom-up dynamic programming approach as Problem 18:

1. Read the triangle from the input file "triangle.txt"
2. Work from the bottom row upward
3. For each element in the current row, add the maximum of the two adjacent elements from the row below
4. Continue this process until reaching the top of the triangle
5. The single value at the top will be the maximum path sum

Specifically, the algorithm:
1. Reads the triangle data from a file, parsing each line into a row of integers
2. Iteratively reduces the triangle by combining rows from bottom to top
3. For each pair of adjacent rows:
   - Create a new row where each element is the sum of the element in the upper row plus the maximum of the two adjacent elements in the lower row
   - Replace the upper row with this new row
   - Remove the lower row
4. Continue until only one row (with one element) remains
5. Return this element as the maximum path sum

This approach has O(n²) time complexity, where n is the number of rows in the triangle, as we process each element in the triangle exactly once. This makes it efficient even for very large triangles, unlike the exponential time complexity of a brute-force approach.