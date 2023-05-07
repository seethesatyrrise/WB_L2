package main

import (
	"os"
)

// вывод текущей директории
func pwd() (interface{}, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	print(wd, "\n")

	return wd, nil
}

// смена директории
func cd(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}

	print("path was changed to ", path, "\n")

	return nil
}
