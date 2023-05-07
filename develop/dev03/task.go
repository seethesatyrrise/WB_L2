package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("type command in format: sort filename [options]...\n" +
		"example: sort ls -r -u -b -c\n" +
		"file must be in folder \"texts\"\n" +
		"type \"sort --help\" for options list")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		com := scanner.Text()

		filename, parameters, needHelp, err := parseCommand(com)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if needHelp {
			getHelp()
			continue
		}

		t := &text{}
		t.lines, err = getLines(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = t.sortFile(parameters)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = t.saveInFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
