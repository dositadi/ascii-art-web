package processor

import (
	m "ascii-web/pkg/models"
	h "ascii-web/pkg/utils"
)

// This specifies the banner the user choose
func (p *Processor) ProcessTextIntoAscii(input m.Ascii) (string, *m.Error) {
	text := input.Text
	banner := input.Banner

	words := h.CheckIfNewlineAndSplit(text)

	var maxStore [][][]string
	var err *m.Error

	switch banner {
	case "standard":
		path := "asset/standard.txt"
		maxStore, err = p.DesignAllCharacters(words, path)
		if err != nil {
			return "", err
		}
	case "shadow":
		path := "asset/shadow.txt"
		maxStore, err = p.DesignAllCharacters(words, path)
		if err != nil {
			return "", err
		}
	case "thinkertoy":
		path := "asset/thinker_toy.txt"
		maxStore, err = p.DesignAllCharacters(words, path)
		if err != nil {
			return "", err
		}
	}

	outPut, err2 := WriteArtOut(maxStore)
	if err2 != nil {
		return "", err2
	}
	
	return outPut, nil
}
