package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input url (and optionally an output file name) or name of the file with the list of URLs\n" +
		"examples: \"url https://www.google.com/ google\" \"file urls\"")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		c, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
		}
		c.run()
	}
}
