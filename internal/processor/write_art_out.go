package processor

import (
	m "ascii-web/pkg/models"
	h "ascii-web/pkg/utils"
	"slices"
	"strings"
)

func WriteArtOut(maxStore [][][]string) (string, *m.Error) {
	if maxStore == nil {
		return "", &m.Error{
			Error:  h.SERVER_ERR + "3",
			Detail: h.SERVER_ERR_DETAIL,
		}
	}

	var output strings.Builder

	for _, lines := range maxStore {
		currentLines := lines
		lenLines := len(currentLines)
		depth := 8

		var newLineChar [][]string

		newLineChar = append(newLineChar, []string{"\n"})

		if slices.Compare(currentLines[0], newLineChar[0]) == 0 {
			output.WriteString("\n")
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
	return output.String(), nil
}
