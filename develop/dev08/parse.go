package main

import "strings"

// парсинг пайпа
func parsePipe(input string) (*pipe, error) {
	newPipe := &pipe{}
	coms := strings.Split(input, "|")

	for _, com := range coms {
		com = strings.TrimSpace(com)
		parsedCom, err := parseCommand(com)
		if err != nil {
			return nil, err
		}
		newPipe.coms = append(newPipe.coms, *parsedCom)
	}

	return newPipe, nil
}

// парсинг команды
func parseCommand(com string) (*command, error) {
	newCommand := &command{}

	com = strings.TrimSpace(com)
	splitIndex := strings.Index(com, " ")
	if splitIndex == -1 {
		newCommand.com = com
		return newCommand, nil
	}

	newCommand.com = com[:splitIndex]
	newCommand.arg = com[splitIndex+1:]

	return newCommand, nil
}
