package utils

import (
	"bufio"
	"strings"
)

// This function checks if an empty strings occurs in the input and slits
func CheckIfNewlineAndSplit(input string) []string {
	var output []string

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		lines := scanner.Text()
		output = append(output, lines)
	}
	return output
}
