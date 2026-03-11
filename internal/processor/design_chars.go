package processor

import (
	m "ascii-web/pkg/models"
	h "ascii-web/pkg/utils"
	"strings"
	"unicode"
)

// this function constructs the ascii character representation
func (p *Processor) DesignAllCharacters(inputs []string, filepath string) ([][][]string, *m.Error) {
	var maxStore [][][]string

	for i := 0; i < len(inputs); i++ {
		currentInput := inputs[i]
		lines := [][]string{}

		if currentInput != "" {
			for _, rn := range currentInput {
				if unicode.IsLetter(rn) {
					if unicode.IsLower(rn) {
						startLine := h.GetLowerCaseStartLine(rn)
						result, err := p.ReadFileByLine(startLine, filepath)
						if err != nil {
							return nil, err
						}
						lines = append(lines, result)
					} else if unicode.IsUpper(rn) {
						startLine := h.GetUpperCaseStartLine(rn)
						result, err := p.ReadFileByLine(startLine, filepath)
						if err != nil {
							return nil, err
						}
						lines = append(lines, result)
					}
				} else if unicode.IsDigit(rn) {
					startLine := h.GetNumberStartLine(rn)
					result, err := p.ReadFileByLine(startLine, filepath)
					if err != nil {
						return nil, err
					}
					lines = append(lines, result)
				} else if strings.ContainsAny(string(rn), " !\"#$%&()*+'-./:;<=>?@[]\\^`{}|~_,") {
					startLine := h.GetSpecialCharStartLine(rn)
					result, err := p.ReadFileByLine(startLine, filepath)
					if err != nil {
						return nil, err
					}
					lines = append(lines, result)
					continue
				}
			}
			maxStore = append(maxStore, lines)
		} else {
			lines = append(lines, []string{string(rune('\n'))})
			maxStore = append(maxStore, lines)
		}
	}
	return maxStore, nil
}
