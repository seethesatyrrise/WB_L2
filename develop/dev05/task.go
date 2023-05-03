package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input command in format: grep [options]... regexp filename\n" +
		"example: grep -A1 -v -c -n '.*-r-.*' test\n" +
		"\"type \"grep --help\" for list of options\"")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		com := scanner.Text()

		comParsed, needHelp, err := parseCommand(com)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if needHelp {
			getHelp()
			continue
		}

		t, err := getLines(comParsed.filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = t.grepFile(comParsed.parameters, comParsed.reStr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = t.saveInFile(comParsed.filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
