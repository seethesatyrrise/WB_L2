package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func searchMonth(s string, months []string) int {
	for i, month := range months {
		if strings.Contains(s, month) {
			return i
		}
	}
	return -1
}

func getNumber(line string) float64 {
	num := 0.0
	gotPoint := false
	digitsAfterPoint := 1.0
	for _, char := range line {
		if !gotPoint && char == '.' {
			gotPoint = true
		}
		if char >= '0' && char <= '9' {
			if gotPoint {
				num = num + float64(char-'0')/math.Pow(10, digitsAfterPoint)
				digitsAfterPoint++
			} else {
				num = num*10 + float64(char-'0')
			}
		}
	}
	return num
}

func (t *text) reverse() {
	tLen := len(t.lines) - 1
	for i := 0; i <= tLen/2; i++ {
		t.lines[i], t.lines[tLen-i] = t.lines[tLen-i], t.lines[i]
	}
}

func (t *text) isSorted() bool {
	return sort.StringsAreSorted(t.lines)
}

func (t *text) removeDuplicates() {
	m := make(map[string]struct{})
	for i, line := range t.lines {
		if _, found := m[line]; found {
			t.removeLine(i)
		}
		m[line] = struct{}{}
	}
}

func (t *text) removeLine(i int) {
	t.lines = append(t.lines[:i], t.lines[i+1:]...)
}

func (t *text) removeTailSpaces() {
	spaces := 0
	for i, line := range t.lines {
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] == ' ' {
				spaces++
			} else {
				break
			}
		}
		t.lines[i] = t.lines[i][:len(line)-spaces]
	}
}

func getHelp() {
	fmt.Println(
		"-k — указание колонки для сортировки (слова в строке " +
			"могут выступать в качестве колонок, по умолчанию " +
			"разделитель — пробел)\n" +
			"-n — сортировать по числовому значению\n" +
			"-r — сортировать в обратном порядке\n" +
			"-u — не выводить повторяющиеся строки\n" +
			"-M — сортировать по названию месяца\n" +
			"-b — игнорировать хвостовые пробелы\n" +
			"-c — проверять отсортированы ли данные")
}
