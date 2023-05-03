package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		res, err := unpackString(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("result:\n" + res)
	}
}
