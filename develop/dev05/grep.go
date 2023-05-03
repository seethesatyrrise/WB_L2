package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// запуск фильтрации с заданными параметрами
func (t *text) grepFile(parameters map[uint8]int, reStr string) error {
	var err error

	if _, ok := parameters['i']; ok {
		t.re, err = regexp.Compile(strings.ToLower(reStr))
		if err != nil {
			return err
		}
		t.ignoreCase = true
	} else {
		t.re, err = regexp.Compile(reStr)
		if err != nil {
			return err
		}
	}

	if _, ok := parameters['F']; ok {
		t.matches = t.searchFixed(reStr)
	} else {
		t.matches, err = t.searchLines()
		if err != nil {
			return err
		}
	}

	if _, ok := parameters['c']; ok {
		fmt.Println("matches found:", len(t.matches))
	}

	if _, ok := parameters['n']; ok {
		t.printLinesIndexes()
	}

	if num, ok := parameters['C']; ok {
		t.addLines(num, true, true)
	} else {
		if num, ok := parameters['A']; ok {
			t.addLines(num, true, false)
		}
		if num, ok := parameters['B']; ok {
			t.addLines(num, false, true)
		}
	}

	sort.Ints(t.matches)
	t.removeDuplicates()

	if _, ok := parameters['v']; ok {
		t.invertMatches()
	}

	return nil
}
