package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type connection struct {
	conn    net.Conn
	host    string
	port    string
	timeout int
}

func main() {
	fmt.Println("input command in format go-telnet [timeout] host port\n" +
		"examples: \"go-telnet google.com 80\" \"go-telnet --timeout=3s 1.1.1.1 123\"")
	fmt.Print("> ")

	scanner := bufio.NewReader(os.Stdin)
	input, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := parseCommand(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.connect()
	if err != nil {
		fmt.Println("no connection")
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go c.write(wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-signalChan
	fmt.Println("closing connection")
	wg.Wait()
	c.closeConnection()
}
