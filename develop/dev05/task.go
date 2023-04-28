package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type command struct {
	filename   string
	parameters map[uint8]int
	reStr      string
}

type text struct {
	lines      []string
	matches    []int
	re         *regexp.Regexp
	ignoreCase bool `default:"false"`
}

func main() {
	fmt.Println("input command in format: grep [options]... regexp filename\n" +
		"example: grep -A1 -v -c -n '.*-r-.*' test")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		com := scanner.Text()
		err := grepFile(com)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// запуск фильтрации
func grepFile(com string) error {
	comParsed, err := parseCommand(com)
	if err != nil {
		return err
	}

	t := &text{}
	if _, ok := comParsed.parameters['i']; ok {
		t.re, err = regexp.Compile(strings.ToLower(comParsed.reStr))
		if err != nil {
			return err
		}
		t.ignoreCase = true
	} else {
		t.re, err = regexp.Compile(comParsed.reStr)
		if err != nil {
			return err
		}
	}

	err = t.getLines(comParsed.filename)
	if err != nil {
		return err
	}

	if _, ok := comParsed.parameters['F']; ok {
		t.matches = t.searchFixed(comParsed.reStr)
	} else {
		t.matches, err = t.searchLines()
		if err != nil {
			return err
		}
	}

	if _, ok := comParsed.parameters['c']; ok {
		fmt.Println("matches found:", len(t.matches))
	}

	if _, ok := comParsed.parameters['n']; ok {
		t.printLinesIndexes()
	}

	if num, ok := comParsed.parameters['C']; ok {
		t.addLines(num, true, true)
	} else {
		if num, ok := comParsed.parameters['A']; ok {
			t.addLines(num, true, false)
		}
		if num, ok := comParsed.parameters['B']; ok {
			t.addLines(num, false, true)
		}
	}

	sort.Ints(t.matches)
	t.removeDuplicates()

	if _, ok := comParsed.parameters['v']; ok {
		t.invertMatches()
	}

	err = t.saveInFile(comParsed.filename)
	if err != nil {
		return err
	}

	return nil
}

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

func (t *text) printLinesIndexes() {
	fmt.Print("matches found in lines: ")
	for _, index := range t.matches {
		fmt.Print(index, " ")
	}
	fmt.Println()
}

func parseCommand(com string) (*command, error) {
	w := strings.Fields(com)
	if w[0] != "grep" {
		return nil, errors.New("undefined command")
	}
	params := make(map[uint8]int)
	if len(w) > 3 {
		params = getParameters(w[1 : len(w)-2])
	}
	re := w[len(w)-2]
	return &command{
		filename:   w[len(w)-1],
		parameters: params,
		reStr:      re[1 : len(re)-1],
	}, nil
}

func (t *text) getLines(name string) error {
	file, err := os.Open("develop\\dev05\\" + name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if t.ignoreCase {
			line = strings.ToLower(line)
		}
		t.lines = append(t.lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return err
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

func (t *text) saveInFile(name string) error {
	file, err := os.Create("develop\\dev05\\grep_" + name)
	defer file.Close()
	if err != nil {
		return err
	}

	tail := "\n"
	for i, lineNum := range t.matches {
		if i == len(t.matches)-1 {
			tail = ""
		}
		_, err = file.Write([]byte(t.lines[lineNum] + tail))
		if err != nil {
			return err
		}
	}

	fmt.Println("result saved in file grep_" + name)
	return nil
}
