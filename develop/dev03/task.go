package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type command struct {
	filename   string
	parameters map[uint8]int
}

func main() {
	com := "sort text.txt -k -u"
	res, err := parseCommand(com)
	fmt.Println(res, err)
	//name := "text.txt"
	//text, _ := getLines(name)
	//sortLines(text)
	//saveInFile(text, name)
}

func parseCommand(com string) (*command, error) {
	w := strings.Fields(com)
	if w[0] != "sort" {
		return nil, errors.New("undefined command")
	}
	return &command{
		filename:   w[1],
		parameters: getParameters(w[2:]),
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

func sortLines(text []string) {
	sort.Strings(text)
}

func saveInFile(text []string, name string) {
	file, err := os.Create("develop\\dev03\\sorted_" + name)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range text {
		file.Write([]byte(line + "\n"))
	}
}

func getParameters(in []string) map[uint8]int {
	params := make(map[uint8]int)
	for _, p := range in {
		if p[0] != '-' {
			continue
		}
		num, err := strconv.Atoi(p[2:])
		if err != nil {
			num = 0
		}
		params[p[1]] = num
	}
	return params
}
