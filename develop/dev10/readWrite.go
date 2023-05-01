package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

func (c *connection) write(wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		query, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("closing writer")
			break
		}
		_, err = c.conn.Write([]byte(fmt.Sprintf("%s\r\n", query)))
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.read()
	}
}

func (c *connection) read() {
	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 256)
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
	fmt.Println("\n" + string(buf))
}
