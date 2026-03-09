package utils

import (
	"fmt"
	"slices"
	"strings"
)

func PrintArtOut(maxStore [][][]string) string {
	if maxStore == nil {
		fmt.Println("The Max Store is empty.")
		return ""
	}

	var output strings.Builder

	for _, lines := range maxStore {
		currentLines := lines
		lenLines := len(currentLines)
		depth := 8

		var newLineChar [][]string

		newLineChar = append(newLineChar, []string{"\n"})

		if slices.Compare(currentLines[0], newLineChar[0]) == 0 {
			fmt.Println()
			continue
		}

		for i := 0; i < depth; i++ {
			combinedString := ""
			for j := 0; j < lenLines; j++ {
				if j != lenLines-1 {
					combinedString += currentLines[j][i] + " "
				} else {
					combinedString += currentLines[j][i] + "\n"
				}
			}
			output.WriteString(combinedString)
		}
	}
	return output.String()
}
