package main

import (
	"errors"
	"strconv"
	"strings"
)

// парсинг команды
func parseCommand(com string) (string, map[uint8]int, bool, error) {
	w := strings.Fields(com)

	if w[0] != "sort" {
		return "", nil, false, errors.New("undefined command")
	}

	if len(w) <= 1 {
		return "", nil, false, errors.New("no filename")
	}

	if w[1] == "--help" {
		return "", nil, true, nil
	}

	var parameters map[uint8]int
	if len(w) > 2 {
		parameters = getParameters(w[2:])
	} else {
		parameters = make(map[uint8]int)
	}

	return w[1], parameters, false, nil
}

// парсинг параметров
func getParameters(in []string) map[uint8]int {
	parameters := make(map[uint8]int)

	for _, p := range in {
		if p[0] != '-' {
			continue
		}
		if len(p) > 2 {
			num, err := strconv.Atoi(p[2:])
			if err != nil {
				num = 0
			}
			parameters[p[1]] = num
			continue
		}
		parameters[p[1]] = 0
	}
	return parameters
}
