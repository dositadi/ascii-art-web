package utils

func GetLowerCaseStartLine(char rune) int {
	if lineNumber, ok := LowerCaseAlphabetsArt[char]; ok {
		return lineNumber
	}
	return -1
}
