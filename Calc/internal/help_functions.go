package calc

import (
	"unicode/utf8"
)

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func isOperation(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func isOpenedBracket(char rune) bool {
	return char == '('
}

func isClosedBracket(char rune) bool {
	return char == ')'
}

func isEndline(char rune) bool {
	return int(char) == 10 || int(char) == 13
}

func getSize(str string) int {
	size := utf8.RuneCountInString(str)
	if size > 0 && isEndline(rune(str[size-1])) {
		size--
	}
	return size
}
