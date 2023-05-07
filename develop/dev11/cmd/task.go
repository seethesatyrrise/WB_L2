package main

import (
	"WB-L2/develop/dev11/internal/rest"
	"WB-L2/develop/dev11/internal/server"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errCode := 0
	defer func() {
		os.Exit(errCode)
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := rest.NewCalendar()
	config := server.GetConfig()

	s := server.NewServer(c, config)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	if err := server.Start(s); err != nil {
		fmt.Println("can't start server")
		errCode = 1
		return
	}

	<-signalChan
	if err := server.Shutdown(ctx, s); err != nil {
		fmt.Println("error shutting server")
		errCode = 1
		return
	}
}
