package main

import (
	"errors"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

// вывод всех процессов
func processes() error {
	processList, err := ps.Processes()
	if err != nil {
		return err
	}

	for _, process := range processList {
		print(process.Pid(), "\t", process.Executable(), "\n")
	}

	return nil
}

// убить процесс
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

	print("process ", proc.Pid, " was killed", "\n")
	return nil
}

// windows doesn't support fork
func fork() (interface{}, error) {
	args := os.Args
	childProc, err := os.StartProcess(args[0], args, &os.ProcAttr{})
	if err != nil {
		return 0, err
	}

	print("child process id: ", childProc.Pid, "\n")

	return childProc.Pid, nil
}

// windows doesn't support exec
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

	print("new process id: ", newProc.Pid, "\n")

	return newProc.Pid, nil
}
