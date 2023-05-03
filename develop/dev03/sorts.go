package main

import (
	"fmt"
	"sort"
	"strings"
)

func (t *text) sortFile(parameters map[uint8]int) error {

	if _, ok := parameters['b']; ok {
		t.removeTailSpaces()
	}

	if _, ok := parameters['u']; ok {
		t.removeDuplicates()
	}

	if _, ok := parameters['c']; ok {
		if t.isSorted() {
			fmt.Println("file is sorted")
		} else {
			fmt.Println("file is not sorted")
		}
	}

	if _, ok := parameters['n']; ok {
		t.sortNumerically()
	}

	if k, ok := parameters['k']; ok {
		t.sortByColumn(k)
	} else {
		t.sortLines()
	}

	if _, ok := parameters['M']; ok {
		t.sortByMonth()
	}

	if _, ok := parameters['r']; ok {
		t.reverse()
	}

	return nil
}

func (t *text) sortLines() {
	sort.Strings(t.lines)
}

func (t *text) sortByColumn(k int) {
	sort.Slice(t.lines, func(i, j int) bool {
		iWords := strings.Fields(t.lines[i])
		jWords := strings.Fields(t.lines[j])
		if len(iWords) <= k {
			return true
		}
		if len(jWords) <= k {
			return false
		}
		fmt.Println(iWords[k], jWords[k])
		return iWords[k] < jWords[k]
	})
}

func (t *text) sortByMonth() {
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	lineMonths := make([]int, len(t.lines))
	for i, line := range t.lines {
		lineMonths[i] = searchMonth(line, months)
	}
	sort.Slice(t.lines, func(i, j int) bool {
		iMonth := lineMonths[i]
		if iMonth == -1 {
			return true
		}
		jMonth := lineMonths[j]
		if iMonth == -1 {
			return false
		}
		return iMonth < jMonth
	})
}

func (t *text) sortNumerically() {
	nums := make([]float64, len(t.lines))
	for i, line := range t.lines {
		nums[i] = getNumber(line)
	}
	sort.Slice(t.lines, func(i, j int) bool {
		iNum := nums[i]
		if iNum == 0 {
			return true
		}
		jNum := nums[j]
		if jNum == 0 {
			return false
		}
		return iNum < jNum
	})
}
