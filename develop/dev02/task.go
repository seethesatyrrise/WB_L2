package main

import (
	"errors"
	"fmt"
)

func main() {
	// "a4bc2d5e" => "aaaabccddddde"

	res, err := unpackString("qwe\\\\5")
	fmt.Println(res, err)
}

func unpackString(s string) (string, error) {
	res := []rune{}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if isDigit(r) {
			if i == 0 {
				return "", errors.New("invalid string")
			}
			continue
		}
		if isBackslash(r) {
			if i == len(runes)-1 {
				return "", errors.New("invalid string")
			}
			i++
			r = runes[i]
		}
		res = append(res, unpackRune(r, runes[i+1:])...)
	}
	return string(res), nil
}

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

func unpackRune(r rune, str []rune) []rune {
	num := 0
	res := []rune{}
	for _, char := range str {
		if isDigit(char) {
			num = num*10 + int(char-'0')
		} else {
			break
		}
	}
	if num == 0 {
		num = 1
	}
	for i := 0; i < num; i++ {
		res = append(res, r)
	}
	return res
}
