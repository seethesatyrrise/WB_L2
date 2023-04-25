package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type command struct {
	filename   string
	parameters map[uint8]int
}

type text struct {
	lines []string
}

func main() {
	com := "sort ls -M"
	err := sortFile(com)
	fmt.Println(err)
}

func sortFile(com string) error {
	comParsed, err := parseCommand(com)
	if err != nil {
		return err
	}

	t := &text{}
	t.lines, err = getLines(comParsed.filename)
	if err != nil {
		return err
	}

	if _, ok := comParsed.parameters['b']; ok {
		t.removeTailSpaces()
	}

	if _, ok := comParsed.parameters['u']; ok {
		t.removeDuplicates()
	}

	if _, ok := comParsed.parameters['c']; ok {
		if t.isSorted() {
			fmt.Println("file is sorted")
		} else {
			fmt.Println("file is not sorted")
		}
	}

	if _, ok := comParsed.parameters['n']; ok {
		t.sortNumerically()
	}

	if k, ok := comParsed.parameters['k']; ok {
		t.sortByColumn(k)
	} else {
		t.sortLines()
	}

	if _, ok := comParsed.parameters['M']; ok {
		t.sortByMonth()
	}

	if _, ok := comParsed.parameters['r']; ok {
		t.reverse()
	}

	err = t.saveInFile(comParsed.filename)
	if err != nil {
		return err
	}

	return nil
}

func parseCommand(com string) (*command, error) {
	w := strings.Fields(com)
	if w[0] != "sort" {
		return nil, errors.New("undefined command")
	}
	params := make(map[uint8]int)
	if len(w) > 2 {
		params = getParameters(w[2:])
	}
	return &command{
		filename:   w[1],
		parameters: params,
	}, nil
}

func getLines(name string) ([]string, error) {
	text := []string{}
	file, err := os.Open("develop\\dev03\\" + name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return text, nil
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

func searchMonth(s string, months []string) int {
	for i, month := range months {
		if strings.Contains(s, month) {
			return i
		}
	}
	return -1
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

func (t *text) saveInFile(name string) error {
	file, err := os.Create("develop\\dev03\\sorted_" + name)
	defer file.Close()
	if err != nil {
		return err
	}

	for i, line := range t.lines {
		tail := ""
		if i < len(t.lines)-1 {
			tail = "\n"
		}
		_, err = file.Write([]byte(line + tail))
		if err != nil {
			return err
		}
	}

	return nil
}

func getParameters(in []string) map[uint8]int {
	params := make(map[uint8]int)
	for _, p := range in {
		if p[0] != '-' {
			continue
		}
		if len(p) > 2 {
			num, err := strconv.Atoi(p[2:])
			if err != nil {
				num = 0
			}
			params[p[1]] = num
			continue
		}
		params[p[1]] = 0
	}
	return params
}
