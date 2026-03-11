package utils

// This function gets a lower case letter startline
func GetLowerCaseStartLine(char rune) int {
	if lineNumber, ok := LowerCaseAlphabetsArt[char]; ok {
		return lineNumber
	}
	return -1
}
