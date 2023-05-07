package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit := make(chan struct{})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go run(quit)

	select {
	case <-signalChan:
	case <-quit:
	}
}
