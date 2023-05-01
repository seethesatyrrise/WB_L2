package main

import (
	"errors"
	"strconv"
	"strings"
)

func parseCommand(com string) (*connection, error) {
	newConn := &connection{timeout: 10}

	com = strings.TrimSpace(com)
	words := strings.Fields(com)
	if words[0] != "go-telnet" {
		return nil, errors.New("invalid command")
	}

	if len(words) < 3 {
		return nil, errors.New("too few arguments")
	}

	if len(words[1]) > 11 && words[1][:10] == "--timeout=" {
		if len(words) < 4 {
			return nil, errors.New("too few arguments")
		}
		nStr := words[1][10 : len(words[1])-1]
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return nil, err
		}
		newConn.timeout = n
		newConn.host = words[2]
		newConn.port = words[3]
	} else {
		newConn.host = words[1]
		newConn.port = words[2]
	}

	return newConn, nil
}
