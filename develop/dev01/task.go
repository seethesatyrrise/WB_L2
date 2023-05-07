package main

import (
	"os"
	"task/gettime"
)

func main() {
	errCode := 0
	defer func() {
		os.Exit(errCode)
	}()
	errCode = gettime.GetTime()
}
