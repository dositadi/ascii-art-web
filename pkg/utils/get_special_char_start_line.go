package utils

func GetSpecialCharStartLine(char rune) int {
	if lineNumber, ok := SpecialCharactersArt[char]; ok {
		return lineNumber
	}
	return -1
}
