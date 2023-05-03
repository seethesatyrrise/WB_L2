package main

import (
	"fmt"
	"strings"
)

// добавление строк до/после к найденным совпадениям
func (t *text) addLines(num int, after bool, before bool) {
	added := []int{}

	for _, match := range t.matches {
		if after {
			for i := 1; i <= num && match+i < len(t.lines); i++ {
				added = append(added, match+i)
			}
		}
		if before {
			for i := 1; i <= num && match+i >= 0; i++ {
				added = append(added, match-i)
			}
		}
	}

	t.matches = append(t.matches, added...)
}

// удаление повторяющихся строк
func (t *text) removeDuplicates() {
	prev := -1

	for i := 0; i < len(t.matches); {
		if t.matches[i] != prev {
			prev = t.matches[i]
			i++
		} else {
			if i == len(t.matches) {
				t.matches = t.matches[:i]
			}
			t.matches = append(t.matches[:i], t.matches[i+1:]...)
		}
	}
}

// инверсия результата (вместо совпадения, исключать)
func (t *text) invertMatches() {
	inverted := []int{}

	for i, j := 0, 0; i < len(t.lines); i++ {
		if j < len(t.matches) && i == t.matches[j] {
			j++
		} else {
			inverted = append(inverted, i)
		}
	}

	t.matches = inverted
}

// поиск подходящих под регулярное выражение строк
func (t *text) searchLines() ([]int, error) {
	matches := []int{}

	for i, line := range t.lines {
		found := t.re.MatchString(line)
		if found {
			matches = append(matches, i)
		}
	}

	return matches, nil
}

// поиск строк, содержащих заданную строку
func (t *text) searchFixed(str string) []int {
	matches := []int{}

	for i, line := range t.lines {
		found := strings.Contains(line, str)
		if found {
			matches = append(matches, i)
		}
	}

	return matches
}

// вывод индексов строк с совпадениями
func (t *text) printLinesIndexes() {
	fmt.Print("matches found in lines: ")

	for _, index := range t.matches {
		fmt.Print(index, " ")
	}

	fmt.Println()
}

// вывод описания параметров
func getHelp() {
	fmt.Println(
		"-A - \"after\" печатать +N строк после совпадения\n" +
			"-B - \"before\" печатать +N строк до совпадения\n" +
			"-C - \"context\" (A+B) печатать ±N строк вокруг совпадения\n" +
			"-c - \"count\" (количество строк)\n" +
			"-i - \"ignore-case\" (игнорировать регистр)\n" +
			"-v - \"invert\" (вместо совпадения, исключать)\n" +
			"-F - \"fixed\", точное совпадение со строкой, не паттерн\n" +
			"-n - \"line num\", напечатать номер строки")
}
