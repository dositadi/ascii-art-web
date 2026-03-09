package processor

import (
	m "ascii-web/pkg/models"
	h "ascii-web/pkg/utils"
	"strings"
	"unicode"
)

func (p *Processor) ProcessTextIntoAscii(input m.Ascii) string {
	text := input.Text
	banner := input.Banner

	words := h.CheckIfNewlineAndSplit(text)

	var maxStore [][][]string

	switch banner {
	case "standard":
		path := "/home/gamp/L2E Fellowship/piscine-prompt/ascii-art-web/asset/standard.txt"
		maxStore = p.DesignAllCharacters(words, path)
	case "shadow":
		path := "/home/gamp/L2E Fellowship/piscine-prompt/ascii-art-web/asset/shadow.txt"
		maxStore = p.DesignAllCharacters(words, path)
	case "thinkertoy":
		path := "/home/gamp/L2E Fellowship/piscine-prompt/ascii-art-web/asset/thinker_toy.txt"
		maxStore = p.DesignAllCharacters(words, path)
	}

	outPut := h.PrintArtOut(maxStore)

	return outPut
}

func (p *Processor) DesignAllCharacters(inputs []string, filepath string) [][][]string {
	var maxStore [][][]string

	for i := 0; i < len(inputs); i++ {
		currentInput := inputs[i]
		lines := [][]string{}

		if currentInput != "" {
			for _, rn := range currentInput {
				if unicode.IsLetter(rn) {
					if unicode.IsLower(rn) {
						startLine := h.GetLowerCaseStartLine(rn)
						result, err := h.ReadFileByLine(startLine, filepath)
						if err != nil {
							return nil // Return error
						}
						lines = append(lines, result)
					} else if unicode.IsUpper(rn) {
						startLine := h.GetUpperCaseStartLine(rn)
						result, err := h.ReadFileByLine(startLine, filepath)
						if err != nil {
							return nil
						}
						lines = append(lines, result)
					}
				} else if unicode.IsDigit(rn) {
					startLine := h.GetNumberStartLine(rn)
					result, err := h.ReadFileByLine(startLine, filepath)
					if err != nil {
						return nil
					}
					lines = append(lines, result)
				} else if strings.ContainsAny(string(rn), " !\"#$%&()*+'-./:;<=>?@[]\\^`{}|~_,") {
					startLine := h.GetSpecialCharStartLine(rn)
					result, err := h.ReadFileByLine(startLine, filepath)
					if err != nil {
						return nil
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
	return maxStore
}
