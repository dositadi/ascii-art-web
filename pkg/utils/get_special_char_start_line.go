package utils

// This function get the startline of the special characters
func GetSpecialCharStartLine(char rune) int {
	if lineNumber, ok := SpecialCharactersArt[char]; ok {
		return lineNumber
	}
	return -1
}
