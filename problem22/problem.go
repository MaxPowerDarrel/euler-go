package problem22

import (
	"euler/common"
	"fmt"
	"sort"
	"strings"
)

func Solve() int {
	total := 0
	names, err := readNamesFromFile("names.txt")
	if err != nil {
		panic("unable to read file")
	}
	for i, name := range names {
		nameTotal := 0
		for _, letter := range name {
			nameTotal += calculateNameScore(string(letter))
		}
		total += nameTotal * (i + 1)
	}
	return total
}

// readNamesFromFile reads names from a file and returns them sorted.
func readNamesFromFile(filename string) ([]string, error) {
	lines, err := common.ReadFile(filename, parseNamesFromLine)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	var names []string
	for _, lineNames := range lines {
		names = append(names, lineNames...)
	}

	sort.Strings(names)
	return names, nil
}

// parseNamesFromLine extracts and cleans names from a single line.
func parseNamesFromLine(line string) ([]string, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}

	var names []string

	// Split by comma and process each part
	parts := strings.Split(line, ",")
	for _, part := range parts {
		name := cleanName(part)
		if name != "" {
			names = append(names, name)
		}
	}

	return names, nil
}

// cleanName removes quotes and trims whitespace from a name.
func cleanName(raw string) string {
	name := strings.TrimSpace(raw)
	name = strings.Trim(name, `"'`) // Remove both double and single quotes
	return strings.TrimSpace(name)  // Trim again after quote removal
}

// calculateNameScore computes the alphabetical score for a name.
func calculateNameScore(name string) int {
	score := 0
	for _, char := range name {
		score += charToAlphabetPosition(char)
	}
	return score
}

// charToAlphabetPosition returns the alphabetical position of a character (A=1, B=2, etc.).
func charToAlphabetPosition(char rune) int {
	switch {
	case char >= 'A' && char <= 'Z':
		return int(char - 'A' + 1)
	case char >= 'a' && char <= 'z':
		return int(char - 'a' + 1)
	default:
		return 0
	}
}
