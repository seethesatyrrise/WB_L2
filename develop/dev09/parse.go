package main

import (
	"errors"
	"fmt"
	"strings"
)

func parseInput(com string) (*command, error) {
	c := &command{}
	w := strings.Fields(com)

	switch w[0] {
	case "url":
		if len(w) <= 1 {
			return nil, errors.New("not enough arguments")
		}
		c.com = "url"
		c.url = w[1]
		if len(w) > 2 {
			c.file = w[2]
		}
		if len(w) > 3 {
			fmt.Println("too many names, I will use first of them")
		}
	case "file":
		if len(w) <= 1 {
			return nil, errors.New("not enough arguments")
		}
		c.com = "file"
		c.file = w[1]
		if len(w) > 2 {
			fmt.Println("too many names, I will use first of them")
		}
	default:
		fmt.Println("undefined command")
	}

	return c, nil
}
