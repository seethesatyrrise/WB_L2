package main

import "errors"

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
