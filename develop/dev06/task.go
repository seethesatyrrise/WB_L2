package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type parameters struct {
	fields    []int
	delimiter string
	separated bool
}

type text struct {
	lines      []string
	wordsLines [][]string
}

func main() {

	//com := "cut -d \";\" -s -f 1,2"
	//com := "cut -f 2,3 -d \";\" -s"
	err := cut()
	if err != nil {
		fmt.Println(err)
	}

}

func cut() error {
	fmt.Println("введите команду в формате \"cut [options]...\"")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	com := scanner.Text()

	params, err := parseCommand(com)
	if err != nil {
		return err
	}

	t := &text{}
	t.getLines()

	t.cutLines(params.delimiter)

	t.printLines(params)

	return nil
}

func (t *text) cutLines(d string) {
	t.wordsLines = make([][]string, len(t.lines))
	for i, line := range t.lines {
		t.wordsLines[i] = strings.Split(line, d)
	}
}

func (t *text) getLines() {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		newLine := scanner.Text()
		if newLine == "" {
			break
		}
		t.lines = append(t.lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func parseCommand(com string) (*parameters, error) {
	w := strings.Fields(com)
	if w[0] != "cut" {
		return nil, errors.New("undefined command")
	}
	params := &parameters{delimiter: "\t"}
	if len(w) > 1 {
		params.getParameters(w[1:])
	}
	return params, nil
}

func (params *parameters) getParameters(in []string) *parameters {
	for i := 0; i < len(in); {
		switch in[i] {
		case "-d":
			i++
			if i < len(in) {
				re, err := regexp.Compile("^\".\"$")
				if err != nil {
					fmt.Println(err)
				}
				matched := re.MatchString(in[i])
				if matched {
					params.delimiter = string(in[i][1])
					i++
					continue
				}
			}
			fmt.Println("wrong use of option -d, it will be omitted")
		case "-f":
			i++
			if i < len(in) {
				re, err := regexp.Compile("^([0-9]+,)*[0-9]$")
				if err != nil {
					fmt.Println(err)
				}
				matched := re.MatchString(in[i])
				if matched {
					fieldsStr := strings.Split(in[i], ",")
					params.fields = make([]int, len(fieldsStr))
					for j, field := range fieldsStr {
						params.fields[j], _ = strconv.Atoi(field)
					}
					i++
					continue
				}
			}
			fmt.Println("wrong use of option -f, it will be omitted")
		case "-s":
			params.separated = true
			i++
		default:
			i++
		}
	}

	return params
}

func (t *text) printLines(params *parameters) {
	for _, wordsLine := range t.wordsLines {
		if params.separated && len(wordsLine) < 2 {
			continue
		}
		if len(params.fields) == 0 {
			for _, word := range wordsLine {
				fmt.Print(word + "\t")
			}
		} else {
			for _, lineNum := range params.fields {
				if lineNum < 1 || lineNum > len(wordsLine) {
					continue
				}
				fmt.Print(wordsLine[lineNum-1] + "\t")
			}
		}
		fmt.Println()
	}
}
