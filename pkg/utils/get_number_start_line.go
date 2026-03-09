package utils

func GetNumberStartLine(char rune) int {
	if lineNumber, ok := NumbersArt[char]; ok {
		return lineNumber
	}
	return -1
}
