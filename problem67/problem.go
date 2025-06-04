package problem67

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() int {
	triangle, err := readTriangleFromFile("triangle.txt")
	if err != nil {
		panic("unable to read file")
	}

	for len(triangle) > 1 {
		lastIdx := len(triangle) - 1

		newRow, err := combineRows(triangle[lastIdx-1], triangle[lastIdx])
		if err != nil {
			panic(err)
		}

		// Replace the second-to-last row and remove the last row
		triangle[lastIdx-1] = newRow
		triangle = triangle[:lastIdx]
	}

	return triangle[0][0]
}

func readTriangleFromFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var triangle [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		// Split the line by spaces and convert to integers
		fields := strings.Fields(line)
		row := make([]int, len(fields))

		for i, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("invalid number '%s' on line %d", field, len(triangle)+1)
			}
			row[i] = num
		}
		triangle = append(triangle, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return triangle, nil
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
