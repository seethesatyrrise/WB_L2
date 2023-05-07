package main

import (
	"bufio"
	"errors"
	"os"
)

type command struct {
	com string
	arg interface{}
}

type pipe struct {
	coms []command
	quit chan struct{}
}

// чтение команд, выполнение
func run(quit chan struct{}) {
	print("type \"help\" for commands list\n")
	for {
		print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		com := scanner.Text()

		p, err := parsePipe(com)
		p.quit = quit
		if err != nil {
			print(err, "\n")
			continue
		}

		err = p.runPipe()
		if err != nil {
			print(err, "\n")
		}
	}
}

// выполнение пайпа
func (p *pipe) runPipe() error {
	for i := 0; i < len(p.coms); i++ {
		res, err := p.executeCommand(i)
		if err != nil {
			return err
		}
		if i < len(p.coms)-1 {
			p.coms[i+1].arg = res
		}
	}

	return nil
}

// выполнение команды
func (p *pipe) executeCommand(i int) (interface{}, error) {
	c := p.coms[i]

	switch c.com {
	case "pwd":
		return pwd()
	case "cd":
		if c.arg == "" {
			return nil, errors.New("need arg for cd")
		}
		stringArg, ok := c.arg.(string)
		if !ok {
			return nil, errors.New("invalid argument for ch")
		}
		return nil, cd(stringArg)
	case "echo":
		return echo(c.arg), nil
	case "kill":
		if c.arg == "" {
			return nil, errors.New("need arg for kill")
		}
		return nil, kill(c.arg)
	case "ps":
		return nil, processes()
	case "fork":
		return fork()
	case "exec":
		if c.arg == "" {
			return nil, errors.New("need arg for exec")
		}
		return exec(c.arg)
	case "help":
		getHelp()
		return nil, nil
	case "\\quit":
		p.quit <- struct{}{}
		return nil, nil
	default:
	}
	return nil, errors.New("undefined command")
}
