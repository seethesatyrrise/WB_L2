package main

import "fmt"

func print(output ...interface{}) {
	for _, outs := range output {
		fmt.Print(outs)
	}
}
