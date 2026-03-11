package processor

import (
	m "ascii-web/pkg/models"
	h "ascii-web/pkg/utils"
	"bufio"
	"os"
)

// The function reads the file to get the ascii representation of the input
func (p *Processor) ReadFileByLine(startLine int, filepath string) ([]string, *m.Error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, &m.Error{
			Error:  h.SERVER_ERR,
			Detail: h.SERVER_ERR_DETAIL,
		}
	}

	defer file.Close()

	var asciiArt []string

	scanner := bufio.NewScanner(file)

	currentLine := 1

	stopLine := startLine + 8

	for scanner.Scan() {
		if currentLine >= startLine && currentLine <= stopLine {
			asciiArt = append(asciiArt, scanner.Text())
		} else if currentLine > stopLine {
			break
		}
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return nil, &m.Error{
			Error:  h.SERVER_ERR,
			Detail: err.Error(),
		}
	}
	return asciiArt, nil
}
