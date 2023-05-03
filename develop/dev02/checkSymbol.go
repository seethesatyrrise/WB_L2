package main

func isDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func isBackslash(r rune) bool {
	if r == '\\' {
		return true
	}
	return false
}
