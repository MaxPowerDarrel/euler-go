package problem18

import "fmt"

func Solve(input [][]int) int {
	for len(input) > 1 {
		lastIdx := len(input) - 1

		newRow, err := combineRows(input[lastIdx-1], input[lastIdx])
		if err != nil {
			panic(err)
		}

		// Replace the second-to-last row and remove the last row
		input[lastIdx-1] = newRow
		input = input[:lastIdx]
	}

	return input[0][0]
}

func combineRows(topRow, bottomRow []int) ([]int, error) {
	if len(bottomRow) != len(topRow)+1 {
		return nil, fmt.Errorf("top row must be exactly 1 shorter than bottom row")
	}

	newRow := make([]int, len(topRow))
	for i := range topRow {
		left := topRow[i] + bottomRow[i]
		right := topRow[i] + bottomRow[i+1]
		if left > right {
			newRow[i] = left
		} else {
			newRow[i] = right
		}
	}
	return newRow, nil
}
