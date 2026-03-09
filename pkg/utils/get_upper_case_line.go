package utils

func GetUpperCaseStartLine(char rune) int {
	if lineNumber, ok := UpperCaseAlphabetsArt[char]; ok {
		return lineNumber
	}
	return -1
}
