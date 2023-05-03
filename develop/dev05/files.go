package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type text struct {
	lines      []string
	matches    []int
	re         *regexp.Regexp
	ignoreCase bool
}

// чтение строк из файла
func getLines(name string) (*text, error) {
	t := &text{}

	file, err := os.Open("develop\\dev05\\" + name)
	if err != nil {
		log.Fatal(err)
		return nil, err
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
		return nil, err
	}
	return t, nil
}

// сохранение строк в файл
func (t *text) saveInFile(name string) error {
	file, err := os.Create("develop\\dev05\\grep_" + name)
	if err != nil {
		return err
	}
	defer file.Close()

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
