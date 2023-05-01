package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
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
	}

	c, err := parseCommand(input)
	c.connect()

	ctx, cancel := context.WithCancel(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	go c.write(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-signalChan
	fmt.Println("closing connection")
	cancel()
	c.closeConnection()
}

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
		nStr := words[1][11 : len(words[1])-2]
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

func (c *connection) connect() error {
	var err error
	c.conn, err = net.DialTimeout("tcp", c.host+":"+c.port, time.Duration(c.timeout)*time.Second)
	if err != nil {
		return err
	}
	fmt.Println("connected. Input queries:")

	return nil
}

func (c *connection) write(ctx context.Context) {
	//n, err := fmt.Fprintf(c.conn, "GET / HTTP/1.0\r\n\r\n")
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		query, err := reader.ReadString('\n')
		select {
		case <-ctx.Done():
			fmt.Println("closing writer")
			return
		default:
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		n, err := c.conn.Write([]byte(fmt.Sprintf("%s\r\n", query)))
		fmt.Println(n, err)
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.read()
	}
}

func (c *connection) read() {
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := c.conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		buf = append(buf, tmp[:n]...)

	}
	fmt.Println(string(buf))

}

func (c *connection) closeConnection() {
	c.conn.Close()
}
