package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type text struct {
	lines []string
}

func getLines(name string) ([]string, error) {
	text := []string{}
	file, err := os.Open("develop\\dev03\\texts\\" + name)
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

func (t *text) saveInFile(name string) error {
	file, err := os.Create("develop\\dev03\\texts\\sorted_" + name)
	defer file.Close()
	if err != nil {
		return err
	}

	for i, line := range t.lines {
		tail := ""
		if i < len(t.lines)-1 {
			tail = "\n"
		}
		_, err = file.Write([]byte(line + tail))
		if err != nil {
			return err
		}
	}

	fmt.Println("result saved in file sorted_" + name)
	return nil
}
