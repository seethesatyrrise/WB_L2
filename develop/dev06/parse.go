package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type parameters struct {
	fields    []int
	delimiter string
	separated bool
}

// парсинг команды
func parseCommand(com string) (*parameters, bool, error) {
	if com == "" {
		return nil, false, errors.New("empty command")
	}
	w := strings.Fields(com)
	if w[0] != "cut" {
		return nil, false, errors.New("undefined command")
	}

	if len(w) > 1 && w[1] == "--help" {
		return nil, true, nil
	}

	return getParameters(w), false, nil
}

// парсинг параметров
func getParameters(in []string) *parameters {
	params := &parameters{delimiter: "\t"}

	for i := 1; i < len(in); {
		switch in[i] {
		case "-d":
			i++
			if i < len(in) {
				re, err := regexp.Compile("^\".\"$")
				if err != nil {
					fmt.Println(err)
				}
				matched := re.MatchString(in[i])
				if matched {
					params.delimiter = string(in[i][1])
					i++
					continue
				}
			}
			fmt.Println("wrong use of option -d, it will be omitted")
		case "-f":
			i++
			if i < len(in) {
				re, err := regexp.Compile("^([0-9]+,)*[0-9]$")
				if err != nil {
					fmt.Println(err)
				}
				matched := re.MatchString(in[i])
				if matched {
					fieldsStr := strings.Split(in[i], ",")
					params.fields = make([]int, len(fieldsStr))
					for j, field := range fieldsStr {
						params.fields[j], _ = strconv.Atoi(field)
					}
					i++
					continue
				}
			}
			fmt.Println("wrong use of option -f, it will be omitted")
		case "-s":
			params.separated = true
			i++
		default:
			i++
		}
	}

	return params
}
