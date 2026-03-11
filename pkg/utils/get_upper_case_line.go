package utils

// This function gets the startline of an uppercase letter
func GetUpperCaseStartLine(char rune) int {
	if lineNumber, ok := UpperCaseAlphabetsArt[char]; ok {
		return lineNumber
	}
	return -1
}




  