package main

import (
	"fmt"
	"net"
	"time"
)

func (c *connection) connect() error {
	var err error
	c.conn, err = net.DialTimeout("tcp", c.host+":"+c.port, time.Duration(c.timeout)*time.Second)
	if err != nil {
		return err
	}
	fmt.Println("connected. Input queries:")

	return nil
}

func (c *connection) closeConnection() {
	c.conn.Close()
}
