package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// чтение команды, чтение строк, выполнение команды и вывод строк
func cut() {
	fmt.Println("type command in format: cut [options]...\n" +
		"example: cut -f 2,3 -d \";\" -s\n" +
		"type \"cut --help\" for options list")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		com := scanner.Text()
		params, needHelp, err := parseCommand(com)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if needHelp {
			getHelp()
			continue
		}

		t, err := getLines()

		t.cutLines(params.delimiter)

		t.printLines(params)

		if err != nil {
			fmt.Println(err)
		}
	}
}

// разделение строк на столбцы
func (t *text) cutLines(d string) {
	t.wordsLines = make([][]string, len(t.lines))

	for i, line := range t.lines {
		t.wordsLines[i] = strings.Split(line, d)
	}
}
