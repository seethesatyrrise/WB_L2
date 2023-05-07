package main

import (
	"bufio"
	"fmt"
	"os"
)

type text struct {
	lines      []string
	wordsLines [][]string
}

// чтение строк из stdin
func getLines() (*text, error) {
	t := &text{}
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
		return nil, err
	}

	return t, nil
}

// вывод строк в stdout
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
