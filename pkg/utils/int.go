package utils

func AbsDiff(l, r int) int {
	result := l - r
	if result < 0 {
		result *= -1
	}
	return result
}

func IntParseRune(r rune) int {
	return int(r - '0')
}
