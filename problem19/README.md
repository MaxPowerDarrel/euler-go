# Problem 19: Counting Sundays

## Problem Description
You are given the following information:
- 1 Jan 1900 was a Monday.
- Thirty days has September, April, June and November. All the rest have thirty-one, except February which has twenty-eight, and twenty-nine in leap years.
- A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

## Algorithm
The solution uses a day-tracking approach to count Sundays that fall on the first of each month:

1. Initialize a counter for the total number of Sundays that fall on the first of a month
2. Start with January 1, 1900 as day 2 (Monday, where Sunday is day 1)
3. Iterate through each month of each year from 1900 to 2000
4. For each month:
   - Calculate the day of the week for the first day of the month
   - If the year is 1901 or later and the day is a Sunday (day 1), increment the counter
5. Return the total count

Specifically, the algorithm:
1. Tracks the day of the week using a rolling counter (1-7, where 1 is Sunday)
2. Calculates the offset for each month based on its length:
   - 30-day months (April, June, September, November) have an offset of 2
   - February has an offset of 1 in leap years and 0 in non-leap years
   - 31-day months (January, March, May, July, August, October, December) have an offset of 3
3. Applies leap year rules correctly: a year is a leap year if it's divisible by 4, except for century years which must be divisible by 400
4. Only counts Sundays falling on the first of the month from 1901 to 2000

This approach has O(n) time complexity, where n is the number of months in the date range (12 months × 101 years = 1,212 months), making it very efficient.

Special note, Go's date time library was specifically ignored as the author felt that was cheating