package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

type command struct {
	com string
	arg interface{}
}

type pipe struct {
	coms []command
	quit chan struct{}
}

// windows doesn't support fork/exec
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	quit := make(chan struct{})
	go getCommands(ctx, quit)
	<-quit
}

func getCommands(ctx context.Context, quit chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Print("> ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			com := scanner.Text()
			err := run(com, quit)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func run(com string, quit chan struct{}) error {
	p, err := parsePipe(com)
	p.quit = quit
	if err != nil {
		return err
	}
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
	case "\\quit":
		p.quit <- struct{}{}
		return nil, nil
	default:
		return nil, errors.New("undefined command")
	}
	return nil, nil
}

func pwd() (interface{}, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	fmt.Println(wd)
	return wd, nil
}

func cd(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}
	fmt.Println("path was changed to", path)
	return nil
}

func echo(text interface{}) interface{} {
	fmt.Println(text)
	return text
}

func processes() error {
	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for _, process := range processList {
		fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}
	return nil
}

func kill(idIface interface{}) error {
	pid, ok := idIface.(int)
	var err error
	if !ok {
		pidStr, ok := idIface.(string)
		if !ok {
			return errors.New("invalid argument for kill")
		}
		pid, err = strconv.Atoi(pidStr)
		if err != nil {
			return err
		}
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	err = proc.Kill()
	if err != nil {
		return err
	}

	fmt.Println("process", proc.Pid, "was killed")
	return nil
}

func fork() (interface{}, error) {
	args := os.Args
	childProc, err := os.StartProcess(args[0], args, &os.ProcAttr{})
	if err != nil {
		return 0, err
	}
	fmt.Println("child process id:", childProc.Pid)
	return childProc.Pid, nil
}

func exec(argIface interface{}) (interface{}, error) {
	arg, ok := argIface.(string)
	if !ok {
		return nil, errors.New("invalid argument for kill")
	}

	args := strings.Fields(arg)

	newProc, err := os.StartProcess(args[0], args, &os.ProcAttr{
		Files: []*os.File{
			os.Stdin, os.Stdout, os.Stderr,
		},
	})
	if err != nil {
		return 0, err
	}
	fmt.Println("new process id:", newProc.Pid)
	return newProc.Pid, nil
}
