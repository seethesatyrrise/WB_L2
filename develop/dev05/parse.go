package main

import (
	"errors"
	"strconv"
	"strings"
)

type command struct {
	filename   string
	parameters map[uint8]int
	reStr      string
}

// парсинг команды
func parseCommand(com string) (*command, bool, error) {
	w := strings.Fields(com)
	if w[0] != "grep" {
		return nil, false, errors.New("undefined command")
	}

	if len(w) == 2 && w[1] == "--help" {
		return nil, true, nil
	}

	if len(w) < 3 {
		return nil, false, errors.New("not enough args")
	}

	filename := w[len(w)-1]

	var parameters map[uint8]int
	if len(w) > 3 {
		parameters = getParameters(w[1 : len(w)-2])
	} else {
		parameters = make(map[uint8]int)
	}

	reStr := w[len(w)-2]
	if reStr[0] != '\'' || reStr[len(reStr)-1] != '\'' {
		return nil, false, errors.New("can't find regexp string")
	}
	reStr = reStr[1 : len(reStr)-1]

	return &command{
		filename:   filename,
		parameters: parameters,
		reStr:      reStr,
	}, false, nil
}

// парсинг параметров
func getParameters(in []string) map[uint8]int {
	params := make(map[uint8]int)

	for _, p := range in {
		if p[0] != '-' {
			continue
		}
		if len(p) > 2 {
			num, err := strconv.Atoi(p[2:])
			if err != nil {
				num = 0
			}
			params[p[1]] = num
			continue
		}
		params[p[1]] = 0
	}

	return params
}
