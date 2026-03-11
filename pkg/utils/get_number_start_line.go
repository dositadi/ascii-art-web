package utils


// This function gets the start line of a number
func GetNumberStartLine(char rune) int {
	if lineNumber, ok := NumbersArt[char]; ok {
		return lineNumber
	}
	return -1
}
